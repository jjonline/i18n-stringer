// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// i18n-stringer is a tool to automate the creation of methods that satisfy the fmt.Stringer
// interface. Given the name of a (signed or unsigned) integer type T that has constants
// defined, i18n-stringer will create a new self-contained Go source file implementing
//	func (t T) String() string
// The file is created in the same package and directory as the package that defines T.
// It has helpful defaults designed for use with go generate.
//
// Stringer works best with constants that are consecutive values such as created using iota,
// but creates good code regardless. In the future it might also provide custom support for
// constant sets that are bit patterns.
//
// For example, given this snippet,
//
//	package painkiller
//
//	type Pill int
//
//	const (
//		Placebo Pill = iota
//		Aspirin
//		Ibuprofen
//		Paracetamol
//		Acetaminophen = Paracetamol
//	)
//
// running this command
//
//	i18n-stringer -type=Pill
//
// in the same directory will create the file pill_string.go, in package painkiller,
// containing a definition of
//
//	func (Pill) String() string
//
// That method will translate the value of a Pill constant to the string representation
// of the respective constant name, so that the call fmt.Print(painkiller.Aspirin) will
// print the string "Aspirin".
//
// Typically this process would be run using go generate, like this:
//
//	//go:generate i18n-stringer -type=Pill
//
// If multiple constants have the same value, the lexically first matching name will
// be used (in the example, Acetaminophen will print as "Paracetamol").
//
// With no arguments, it processes the package in the current directory.
// Otherwise, the arguments must name a single directory holding a Go package
// or a set of Go source files that represent a single Go package.
//
// The -type flag accepts a comma-separated list of types so a single run can
// generate methods for multiple types. The default output file is t_string.go,
// where t is the lower-cased name of the first type listed. It can be overridden
// with the -output flag.
package main

import (
	"bytes"
	"flag"
	"fmt"
	"go/ast"
	"go/constant"
	"go/format"
	"go/token"
	"go/types"
	"golang.org/x/tools/go/packages"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"sync"
)

var (
	check         = flag.Bool("check", false, "Check missing or useless key-value pairs in TOML")
	typeNames     = flag.String("type", "", "comma-separated list of type names; must be set")
	output        = flag.String("output", "", "output file name; default srcdir/<type>_i18n_string.go")
	tomlpath      = flag.String("tomlpath", "", "set toml i18n file path; default srcdir/i18n")
	defaultlocale = flag.String("defaultlocale", "", "set default locale name; default naturally sorted first")
	ctxkey        = flag.String("ctxkey", "", "key used by context.Value for get locale; default i18nLocale")
	buildTags     = flag.String("tags", "", "comma-separated list of build tags to apply")
)

// Usage is a replacement usage function for the flags package.
func Usage() {
	_, _ = fmt.Fprintf(os.Stderr, "Usage of i18n-stringer:\n")
	_, _ = fmt.Fprintf(os.Stderr, "\ti18n-stringer [flags] -type T [directory]\n")
	_, _ = fmt.Fprintf(os.Stderr, "\ti18n-stringer [flags] -type T -defaultlocale zh-hk -tomlpath lang files... # Must be a single package\n")
	_, _ = fmt.Fprintf(os.Stderr, "For more information, see:\n")
	_, _ = fmt.Fprintf(os.Stderr, "\thttps://github.com/tvb-sz/i18n-stringer\n")
	_, _ = fmt.Fprintf(os.Stderr, "Flags:\n")
	flag.PrintDefaults()
}

func main() {
	log.SetFlags(0)
	log.SetPrefix("i18n-stringer: ")
	flag.Usage = Usage
	flag.Parse()
	if len(*typeNames) == 0 {
		flag.Usage()
		os.Exit(2)
	}
	typeItems := strings.Split(*typeNames, ",")
	var tags []string
	if len(*buildTags) > 0 {
		tags = strings.Split(*buildTags, ",")
	}

	// We accept either one directory or a list of files. Which do we have?
	args := flag.Args()
	if len(args) == 0 {
		// Default: process whole package in current directory.
		args = []string{"."}
	}

	// Parse the package once.
	var dir string
	g := Generator{
		ctxKey:        ternary(*ctxkey, "i18nLocale"),
		tomlPath:      ternary(*tomlpath, "i18n"),
		defaultLocale: ternary(*defaultlocale, ""), // default locale
		values:        make(map[string][]Value),    // init const value
	}

	if len(args) == 1 && isDirectory(args[0]) {
		dir = args[0]
	} else {
		if len(tags) != 0 {
			log.Fatal("-tags option applies only to directories, not when files are specified")
		}
		dir = filepath.Dir(args[0])
	}

	// parse toml locale config file
	g.parser = newParser(g.tomlPath)
	g.parser.parse()

	// set default locale when command param do not set
	if g.defaultLocale == "" {
		g.defaultLocale = g.parser.locales[0] // default naturally sorted first
	}

	// parse package type && const info
	g.parsePackage(args, tags)

	// parse const value for eve Type
	for _, typeName := range typeItems {
		g.parseConstValues(typeName)
	}

	// just check, do not generate, check const and TOML key miss
	if *check {
		g.checkConstDefine()
		os.Exit(0)
	}

	// Print the header and package clause.
	g.Printf("// Code generated by \"i18n-stringer %s\"; DO NOT EDIT.\n", strings.Join(os.Args[1:], " "))
	g.Printf("\n")
	g.Printf("package %s", g.pkg.name)
	g.Printf("\n")
	g.Printf("import (\n")
	g.Printf("\"context\"\n")
	g.Printf("\"strconv\"\n")
	g.Printf(")\n")

	// Run generate for each type.
	for _, typeName := range typeItems {
		g.generate(typeName)
	}

	// Format the output.
	src := g.format()

	// Write to file.
	outputName := *output
	if outputName == "" {
		baseName := fmt.Sprintf("%s_i18n_string.go", typeItems[0])
		outputName = filepath.Join(dir, strings.ToLower(baseName))
	}
	err := os.WriteFile(outputName, src, 0644)
	if err != nil {
		log.Fatalf("writing output: %s", err)
	}
}

// ternary when empty get default
func ternary(from, dVal string) string {
	if from == "" {
		return dVal
	}
	return from
}

// isDirectory reports whether the named file is a directory.
func isDirectory(name string) bool {
	info, err := os.Stat(name)
	if err != nil {
		log.Fatal(err)
	}
	return info.IsDir()
}

// Generator holds the state of the analysis. Primarily used to buffer
// the output for format.Source.
type Generator struct {
	buf           bytes.Buffer       // Accumulated output.
	pkg           *Package           // Package we are scanning.
	parser        *Parser            // toml file Parser
	values        map[string][]Value // parse source code for TYPE CONST values map[typ][]Value
	tomlPath      string
	ctxKey        string
	defaultLocale string
}

func (g *Generator) Printf(format string, args ...interface{}) {
	_, _ = fmt.Fprintf(&g.buf, format, args...)
}

// File holds a single parsed file and associated data.
type File struct {
	pkg  *Package  // Package to which this file belongs.
	file *ast.File // Parsed AST.
	// These fields are reset for each type being generated.
	typeName string  // Name of the constant type.
	values   []Value // Accumulator for constant values of that type.
}

type Package struct {
	name  string
	defs  map[*ast.Ident]types.Object
	files []*File
}

// checkConstDefine check missing CONSTANT
func (g *Generator) checkConstDefine() {
	// The missing key-value pair structure in TOML: map[typ][locale][]K
	var notPairsRecord = make(map[string]map[string][]string)
	for tye, values := range g.values {
		for _, value := range values {
			for locale, items := range g.parser.localesMap {
				if _, exist := items[value.originalName]; !exist {
					if _, existMap := notPairsRecord[tye]; !existMap {
						notPairsRecord[tye] = make(map[string][]string, 0)
					}
					if _, existMap := notPairsRecord[tye][locale]; !existMap {
						notPairsRecord[tye][locale] = make([]string, 0)
					}
					// add do not exist pairs K/V
					notPairsRecord[tye][locale] = append(notPairsRecord[tye][locale], value.originalName)
				}
			}
		}
	}

	// The redundant key-value pairs defined in TOML: map[string][]K
	var noneUsedRecord = make(map[string][]string)
	for locale, items := range g.parser.localesMap {
		for key := range items {
			keyExist := false
			for _, values := range g.values {
				if keyExist {
					break // when key exist, break this key's check
				}
				for _, cValue := range values {
					if cValue.originalName == key {
						keyExist = true
						break
					}
				}
			}

			// all typ do not use this key as CONST
			if !keyExist {
				if _, existMap := noneUsedRecord[locale]; !existMap {
					noneUsedRecord[locale] = make([]string, 0)
				}
				noneUsedRecord[locale] = append(noneUsedRecord[locale], key)
			}
		}
	}

	if len(notPairsRecord) > 0 {
		log.Printf("Check Fail")
		log.Printf("The missing key-value pair information as follows")
		log.Printf("You can copy and fill it to the corresponding TOML file")
		log.SetPrefix("")
		for typ, values := range notPairsRecord {
			for locale, items := range values {
				log.Printf("************TYPE `%s` locale `%s` missing key-value pair************", typ, locale)
				for _, key := range items {
					log.Printf("%s=\"\"", key)
				}
			}
		}
	}

	if len(noneUsedRecord) > 0 {
		log.SetPrefix("i18n-stringer: ")
		log.Printf("Check Warning")
		log.Printf("key-value pairs that will not be used because there is no corresponding defined constant")
		log.Printf("You can delete the key-value pairs in the corresponding TOML file")
		log.SetPrefix("")
		for locale, items := range noneUsedRecord {
			log.Printf("************Can be deleted TOML keys of locale `%s`************", locale)
			for _, key := range items {
				log.Printf("%s", key)
			}
		}
	}

	if len(notPairsRecord) == 0 && len(noneUsedRecord) == 0 {
		log.Printf("Check success, All constants have key-value pairs set")
	}
}

// parsePackage analyzes the single package constructed from the patterns and tags.
// parsePackage exits if there is an error.
func (g *Generator) parsePackage(patterns []string, tags []string) {
	cfg := &packages.Config{
		Mode:       packages.LoadSyntax,
		Tests:      false,
		BuildFlags: []string{fmt.Sprintf("-tags=%s", strings.Join(tags, " "))},
	}
	pkgs, err := packages.Load(cfg, patterns...)
	if err != nil {
		log.Fatal(err)
	}
	if len(pkgs) != 1 {
		log.Fatalf("error: %d packages found", len(pkgs))
	}
	g.addPackage(pkgs[0])
}

// addPackage adds a type checked Package and its syntax files to the generator.
func (g *Generator) addPackage(pkg *packages.Package) {
	g.pkg = &Package{
		name:  pkg.Name,
		defs:  pkg.TypesInfo.Defs,
		files: make([]*File, len(pkg.Syntax)),
	}

	for i, file := range pkg.Syntax {
		g.pkg.files[i] = &File{
			file: file,
			pkg:  g.pkg,
		}
	}
}

// parseConstValues parse const value to g
func (g *Generator) parseConstValues(typeName string) {
	g.values[typeName] = make([]Value, 0, 100)
	for _, file := range g.pkg.files {
		// Set the state for this run of the walker.
		file.typeName = typeName
		file.values = nil
		if file.file != nil {
			ast.Inspect(file.file, file.genDecl)
			g.values[typeName] = append(g.values[typeName], file.values...)
		}
	}

	if len(g.values[typeName]) == 0 {
		log.Fatalf("No values defined in TOML for type %s", typeName)
	}
}

// generate produces the String method for the named type.
func (g *Generator) generate(typeName string) {
	//values := make([]Value, 0, 100)
	//for _, file := range g.pkg.files {
	//	// Set the state for this run of the walker.
	//	file.typeName = typeName
	//	file.values = nil
	//	if file.file != nil {
	//		ast.Inspect(file.file, file.genDecl)
	//		values = append(values, file.values...)
	//	}
	//}
	//
	//if len(values) == 0 {
	//	log.Fatalf("no values defined for type %s", typeName)
	//}

	values := g.values[typeName]

	// Generate code that will fail if the constants change value.
	g.Printf("func _() {\n")
	g.Printf("\t// An \"invalid array index\" compiler error signifies that the constant values have changed.\n")
	g.Printf("\t// Re-run the i18n-stringer command to generate them again.\n")
	g.Printf("\tvar x [1]struct{}\n")
	for _, v := range values {
		g.Printf("\t_ = x[%s - %s]\n", v.originalName, v.str)
	}
	g.Printf("}\n")

	runs := splitIntoRuns(values)

	// The decision of which pattern to use depends on the number of
	// runs in the numbers. If there's only one, it's easy. For more than
	// one, there's a tradeoff between complexity and size of the data
	// and code vs. the simplicity of a map. A map takes more space,
	// but so does the code. The decision here (crossover at 10) is
	// arbitrary, but considers that for large numbers of runs the cost
	// of the linear scan in the switch might become important, and
	// rather than use yet another algorithm such as binary search,
	// we punt and use a map. In any case, the likelihood of a map
	// being necessary for any realistic example other than bitmasks
	// is very low. And bitmasks probably deserve their own analysis,
	// to be done some other day.
	switch {
	case len(runs) == 1:
		g.buildOneRun(runs, typeName)
	case len(runs) <= 10:
		g.buildMultipleRuns(runs, typeName)
	default:
		g.buildMap(runs, typeName)
	}

	// last build common function
	g.buildCommFunc(typeName)
}

// splitIntoRuns breaks the values into runs of contiguous sequences.
// For example, given 1,2,3,5,6,7 it returns {1,2,3},{5,6,7}.
// The input slice is known to be non-empty.
func splitIntoRuns(values []Value) [][]Value {
	// We use stable sort so the lexically first name is chosen for equal elements.
	sort.Stable(byValue(values))
	// Remove duplicates. Stable sort has put the one we want to print first,
	// so use that one. The String method won't care about which named constant
	// was the argument, so the first name for the given value is the only one to keep.
	// We need to do this because identical values would cause the switch or map
	// to fail to compile.
	j := 1
	for i := 1; i < len(values); i++ {
		if values[i].value != values[i-1].value {
			values[j] = values[i]
			j++
		}
	}
	values = values[:j]
	runs := make([][]Value, 0, 10)
	for len(values) > 0 {
		// One contiguous sequence per outer loop.
		i := 1
		for i < len(values) && values[i].value == values[i-1].value+1 {
			i++
		}
		runs = append(runs, values[:i])
		values = values[i:]
	}
	return runs
}

// format returns the gofmt-ed contents of the Generator's buffer.
func (g *Generator) format() []byte {
	src, err := format.Source(g.buf.Bytes())
	if err != nil {
		// Should never happen, but can arise when developing this code.
		// The user can compile the output to see the error.
		log.Printf("warning: internal error: invalid Go generated: %s", err)
		log.Printf("warning: compile the package to analyze the error")
		return g.buf.Bytes()
	}
	return src
}

// Value represents a declared constant.
type Value struct {
	originalName string // The name of the constant.
	name         string // The name with trimmed prefix.
	// The value is stored as a bit pattern alone. The boolean tells us
	// whether to interpret it as an int64 or a uint64; the only place
	// this matters is when sorting.
	// Much of the time the str field is all we need; it is printed
	// by Value.String.
	value  uint64 // Will be converted to int64 when needed.
	signed bool   // Whether the constant is a signed type.
	str    string // The string representation given by the "go/constant" package.
}

func (v *Value) String() string {
	return v.str
}

// byValue lets us sort the constants into increasing order.
// We take care in the Less method to sort in signed or unsigned order,
// as appropriate.
type byValue []Value

func (b byValue) Len() int      { return len(b) }
func (b byValue) Swap(i, j int) { b[i], b[j] = b[j], b[i] }
func (b byValue) Less(i, j int) bool {
	if b[i].signed {
		return int64(b[i].value) < int64(b[j].value)
	}
	return b[i].value < b[j].value
}

// genDecl processes one declaration clause.
func (f *File) genDecl(node ast.Node) bool {
	decl, ok := node.(*ast.GenDecl)
	if !ok || decl.Tok != token.CONST {
		// We only care about const declarations.
		return true
	}
	// The name of the type of the constants we are declaring.
	// Can change if this is a multi-element declaration.
	typ := ""
	// Loop over the elements of the declaration. Each element is a ValueSpec:
	// a list of names possibly followed by a type, possibly followed by values.
	// If the type and value are both missing, we carry down the type (and value,
	// but the "go/types" package takes care of that).
	for _, spec := range decl.Specs {
		vSpec := spec.(*ast.ValueSpec) // Guaranteed to succeed as this is CONST.
		if vSpec.Type == nil && len(vSpec.Values) > 0 {
			// "X = 1". With no type but a value. If the constant is untyped,
			// skip this vSpec and reset the remembered type.
			typ = ""

			// If this is a simple type conversion, remember the type.
			// We don't mind if this is actually a call; a qualified call won't
			// be matched (that will be SelectorExpr, not Ident), and only unusual
			// situations will result in a function call that appears to be
			// a type conversion.
			ce, ok := vSpec.Values[0].(*ast.CallExpr)
			if !ok {
				continue
			}
			id, ok := ce.Fun.(*ast.Ident)
			if !ok {
				continue
			}
			typ = id.Name
		}
		if vSpec.Type != nil {
			// "X T". We have a type. Remember it.
			ident, ok := vSpec.Type.(*ast.Ident)
			if !ok {
				continue
			}
			typ = ident.Name
		}
		if typ != f.typeName {
			// This is not the type we're looking for.
			continue
		}
		// We now have a list of names (from one line of source code) all being
		// declared with the desired type.
		// Grab their names and actual values and store them in f.values.
		for _, name := range vSpec.Names {
			if name.Name == "_" {
				continue
			}
			// This dance lets the type checker find the values for us. It's a
			// bit tricky: look up the object declared by the name, find its
			// types.Const, and extract its value.
			obj, ok := f.pkg.defs[name]
			if !ok {
				log.Fatalf("no value for constant %s", name)
			}
			info := obj.Type().Underlying().(*types.Basic).Info()
			if info&types.IsInteger == 0 {
				log.Fatalf("can't handle non-integer constant type %s", typ)
			}
			value := obj.(*types.Const).Val() // Guaranteed to succeed as this is CONST.
			if value.Kind() != constant.Int {
				log.Fatalf("can't happen: constant is not an integer %s", name)
			}
			i64, isInt := constant.Int64Val(value)
			u64, isUint := constant.Uint64Val(value)
			if !isInt && !isUint {
				log.Fatalf("internal error: value of %s is not an integer: %s", name, value.String())
			}
			if !isInt {
				u64 = uint64(i64)
			}
			v := Value{
				originalName: name.Name,
				value:        u64,
				signed:       info&types.IsUnsigned == 0,
				str:          value.String(),
			}
			v.name = v.originalName
			f.values = append(f.values, v)
		}
	}
	return false
}

// Helpers

// usize returns the number of bits of the smallest unsigned integer
// type that will hold n. Used to create the smallest possible slice of
// integers to use as indexes into the concatenated strings.
func usize(n int) int {
	switch {
	case n < 1<<8:
		return 8
	case n < 1<<16:
		return 16
	default:
		// 2^32 is enough constants for anyone.
		return 32
	}
}

// declareIndexAndNameVars declares the index slices and concatenated names
// strings representing the runs of values.
func (g *Generator) declareIndexAndNameVars(runs [][]Value, typeName string) {
	var indexes, names []string
	for i, run := range runs {
		index, name := g.createIndexAndNameDecl(run, typeName, fmt.Sprintf("_%d", i))
		if len(run) != 1 {
			indexes = append(indexes, index)
		}
		names = append(names, name)
	}
	g.Printf("const (\n")
	for _, name := range names {
		g.Printf("\t%s\n", name)
	}
	g.Printf(")\n\n")

	if len(indexes) > 0 {
		g.Printf("var (")
		for _, index := range indexes {
			g.Printf("\t%s\n", index)
		}
		g.Printf(")\n\n")
	}
}

// declareIndexAndNameVar is the single-run version of declareIndexAndNameVars
func (g *Generator) declareIndexAndNameVar(run []Value, typeName string) {
	index, name := g.createIndexAndNameDecl(run, typeName, "")
	g.Printf("const %s\n", name)
	g.Printf("var %s\n", index)
}

// createIndexAndNameDecl returns the pair of declarations for the run. The caller will add "const" and "var".
func (g *Generator) createIndexAndNameDecl(run []Value, typeName string, suffix string) (string, string) {
	b := new(bytes.Buffer)
	indexes := make([]int, len(run))
	for i := range run {
		b.WriteString(run[i].name)
		indexes[i] = b.Len()
	}
	nameConst := fmt.Sprintf("_%s_name%s = %q", typeName, suffix, b.String())
	nameLen := b.Len()
	b.Reset()
	_, _ = fmt.Fprintf(b, "_%s_index%s = [...]uint%d{0, ", typeName, suffix, usize(nameLen))
	for i, v := range indexes {
		if i > 0 {
			_, _ = fmt.Fprintf(b, ", ")
		}
		_, _ = fmt.Fprintf(b, "%d", v)
	}
	_, _ = fmt.Fprintf(b, "}")
	return b.String(), nameConst
}

// declareNameVars declares the concatenated names string representing all the values in the runs.
func (g *Generator) declareNameVars(runs [][]Value, typeName string, suffix string) {
	g.Printf("const _%s_name%s = \"", typeName, suffix)
	for _, run := range runs {
		for i := range run {
			g.Printf("%s", run[i].name)
		}
	}
	g.Printf("\"\n")
}

// buildOneRun generates the variables and String method for a single run of contiguous values.
func (g *Generator) buildOneRun(runs [][]Value, typeName string) {
	values := runs[0]
	g.Printf("\n")
	g.declareIndexAndNameVar(values, typeName)
	// The generated code is simple enough to write as a Printf format.
	lessThanZero := ""
	if values[0].signed {
		lessThanZero = "i < 0 || "
	}
	if values[0].value == 0 { // Signed or unsigned, 0 is still 0.
		g.Printf(stringOneRun, typeName, usize(len(values)), lessThanZero)
	} else {
		g.Printf(stringOneRunWithOffset, typeName, values[0].String(), usize(len(values)), lessThanZero)
	}
}

// Arguments to format are:
//	[1]: type name
//	[2]: size of index element (8 for uint8 etc.)
//	[3]: less than zero check (for signed types)
const stringOneRun = `// String type to fmt.Stringer interface use default locale
func (i %[1]s) String() string {
	if %[3]si >= %[1]s(len(_%[1]s_index)-1) {
		return "%[1]s(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _%[1]s_name[_%[1]s_index[i]:_%[1]s_index[i+1]]
}
`

// Arguments to format are:
//	[1]: type name
//	[2]: lowest defined value for type, as a string
//	[3]: size of index element (8 for uint8 etc.)
//	[4]: less than zero check (for signed types)
/*
 */
const stringOneRunWithOffset = `// String type to fmt.Stringer interface use default locale
func (i %[1]s) String() string {
	i -= %[2]s
	if %[4]si >= %[1]s(len(_%[1]s_index)-1) {
		return "%[1]s(" + strconv.FormatInt(int64(i + %[2]s), 10) + ")"
	}
	return _%[1]s_name[_%[1]s_index[i] : _%[1]s_index[i+1]]
}
`

// buildMultipleRuns generates the variables and String method for multiple runs of contiguous values.
// For this pattern, a single Printf format won't do.
func (g *Generator) buildMultipleRuns(runs [][]Value, typeName string) {
	g.Printf("\n")
	g.declareIndexAndNameVars(runs, typeName)
	g.Printf("// String type to fmt.Stringer interface use default locale\n")
	g.Printf("func (i %s) String() string {\n", typeName)
	g.Printf("\tswitch {\n")
	for i, values := range runs {
		if len(values) == 1 {
			g.Printf("\tcase i == %s:\n", &values[0])
			g.Printf("\t\treturn _%s_name_%d\n", typeName, i)
			continue
		}
		if values[0].value == 0 && !values[0].signed {
			// For an unsigned lower bound of 0, "0 <= i" would be redundant.
			g.Printf("\tcase i <= %s:\n", &values[len(values)-1])
		} else {
			g.Printf("\tcase %s <= i && i <= %s:\n", &values[0], &values[len(values)-1])
		}
		if values[0].value != 0 {
			g.Printf("\t\ti -= %s\n", &values[0])
		}
		g.Printf("\t\treturn _%s_name_%d[_%s_index_%d[i]:_%s_index_%d[i+1]]\n",
			typeName, i, typeName, i, typeName, i)
	}
	g.Printf("\tdefault:\n")
	g.Printf("\t\treturn \"%s(\" + strconv.FormatInt(int64(i), 10) + \")\"\n", typeName)
	g.Printf("\t}\n")
	g.Printf("}\n")
}

// buildMap handles the case where the space is so sparse a map is a reasonable fallback.
// It's a rare situation but has simple code.
func (g *Generator) buildMap(runs [][]Value, typeName string) {
	g.Printf("\n")
	g.declareNameVars(runs, typeName, "")
	g.Printf("\nvar _%s_map = map[%s]string{\n", typeName, typeName)
	n := 0
	for _, values := range runs {
		for _, value := range values {
			g.Printf("\t%s: _%s_name[%d:%d],\n", &value, typeName, n, n+len(value.name))
			n += len(value.name)
		}
	}
	g.Printf("}\n\n")
	g.Printf(stringMap, typeName)
}

// Argument to format is the type name.
const stringMap = `func (i %[1]s) String() string {
	if str, ok := _%[1]s_map[i]; ok {
		return str
	}
	return "%[1]s(" + strconv.FormatInt(int64(i), 10) + ")"
}
`

// buildCommFunc build common function
func (g *Generator) buildCommFunc(typeName string) {
	g.Printf("\n")
	g.Printf(commFunc, typeName, g.defaultLocale, g.ctxKey)
	g.Printf("\n\n")
}

// Argument to format is the type name.
// 1% typeName
// 2% default locale
// 3% default context get locale key name
const commFunc = `// _%[1]s_defaultLocale default locale generated pass by flag -defaultlocale, Don't assign directly
var _%[1]s_defaultLocale = "%[2]s"

// _%[1]s_supported All supported locales and text offset information
var _%[1]s_supported = map[string]int{"zh-hk":0}

// _%[1]s_ctxKey Key of set used locale from context.Context Value
var _%[1]s_ctxKey = "%[3]s"

func _%[1]s_isLocaleSupport(locale string) bool {
	_, ok := _%[1]s_supported[locale]
	return ok
}

// _%[1]s_localeFromCtxWithFallback retrieves and returns language locale name from context.
// It returns default locale when _%[1]s_isLocaleSupport is false
func _%[1]s_localeFromCtxWithFallback(ctx context.Context) string {
	if ctx == nil {
		return _%[1]s_defaultLocale
	}
	v := ctx.Value(_%[1]s_ctxKey)
	if v != nil && _%[1]s_isLocaleSupport(v.(string)) {
		return v.(string)
	}
	return _%[1]s_defaultLocale
}`

// trans function text
// 1% typeName
const langFunc = `// Lang get target translate text use context.Context
//  - ctx context with Value use Key from  _%[1]s_ctxKey
//  - args Optional placeholder replacement value
func (i _%[1]s) Lang(ctx context.Context, args ...interface{}) string {
    return i._trans(key, _%[1]s_localeFromCtxWithFallback(ctx))
}

// Trans get target translate text use specified language locale identifier
//  - locale specified language locale identifier, need pass by _%[1]s_isLocaleSupport
//  - args Optional placeholder replacement value
func (i _%[1]s) Trans(locale string, args ...interface{}) string {
    if !_%[1]s_isLocaleSupport(locale) {
        locale = _defaultLocale
    }
    return i._trans(locale)
}

// _trans trustworthy parameters inside method
func (i _%[1]s) _trans(locale string, args ...interface{}) string {
    return ""
}`

// +++++++++++++++++++++++++++
// toml file parse util
// +++++++++++++++++++++++++++

// Parser locale config file parser
type Parser struct {
	mu         sync.RWMutex
	files      map[string][]string          // toml file, locale to file list map
	locales    []string                     // naturally sorted, if not specify default locale, first index used
	localesMap map[string]map[string]string // {"locale":{"tran-key": "tran-val", "tran-key1": "tran-val1"}} case-insensitive
	path       string                       // config file belong path
}

// newParser new instance for Parser
func newParser(path string) *Parser {
	if !isDirectory(path) {
		log.Fatal("-tomlpath option applies only to directory, eg. i18n")
	}
	return &Parser{
		files:      make(map[string][]string, 0),
		locales:    make([]string, 0),
		localesMap: make(map[string]map[string]string, 0),
		path:       path,
	}
}

// parse parse toml config file
// toml file just support utf8 K/V mode
//  - CodeErr="aaa"
//  - CodeErr1="aaa\"\n execute"
func (p *Parser) parse() {
	// parse toml file dir list
	dir, err := os.ReadDir(p.path)
	if err != nil {
		log.Fatal(err)
	}
	for _, target := range dir {
		if target.IsDir() {
			locale := target.Name()
			subDir := p.path + "/" + target.Name()
			p.appendTomlFiles(locale, p.listSubDir(subDir))
		} else {
			// just collect .toml suffix file
			if strings.HasSuffix(target.Name(), ".toml") {
				locale := strings.TrimRight(target.Name(), ".toml")
				fileDir := []string{p.path + "/" + target.Name()}
				p.appendTomlFiles(locale, fileDir)
			} else {
				log.Printf("Use only TOML format files, `%s` is ignored\n", p.path+"/"+target.Name())
			}
		}
	}

	// notice if toml none
	if len(p.files) <= 0 {
		log.Fatalf("No valid TOML file found, please write lacale TOML file at first")
		return
	}

	for locale := range p.files {
		p.locales = append(p.locales, locale)
	}

	// naturally sorted
	sort.Sort(sort.StringSlice(p.locales))

	// parse then read toml file K/V
	p.readToml2KV()
}

// appendTomlFiles collect toml file with locale
func (p *Parser) appendTomlFiles(locale string, files []string) {
	if len(files) <= 0 {
		return
	}
	p.mu.Lock()
	defer p.mu.Unlock()
	if _, ok := p.files[locale]; ok {
		p.files[locale] = append(p.files[locale], files...)
	} else {
		p.files[locale] = files
	}
}

// listSubDir list toml sub dir
func (p *Parser) listSubDir(subPath string) []string {
	res := make([]string, 0)
	_ = filepath.WalkDir(subPath, func(path string, d fs.DirEntry, err error) error {
		// just collect .toml suffix file
		if err == nil && !d.IsDir() {
			if strings.HasSuffix(path, ".toml") {
				res = append(res, path)
			} else {
				log.Printf("Use only TOML format files, `%s` is ignored\n", path)
			}
		}
		return nil
	})
	return res
}

// readToml2KV read all toml file to K/V
func (p *Parser) readToml2KV() {
	for locale, files := range p.files {
		for _, file := range files {
			p.readOneToml(file, locale)
		}
	}
}

// readOneToml read one toml file
func (p *Parser) readOneToml(path, locale string) {
	stream, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("read TOML file `%s` occur err %s", path, err.Error())
	}

	// Read over BOM
	data := string(stream)
	if strings.HasPrefix(data, "\xff\xfe") || strings.HasPrefix(data, "\xfe\xff") {
		data = data[2:]
	}

	// TOML files must be UTF-8
	ex := 6
	if len(data) < 6 {
		ex = len(data)
	}
	if i := strings.IndexRune(data[:ex], 0); i > -1 {
		log.Fatalf("TOML file `%s` must be using UTF-8 coding", path)
	}

	lines := strings.Split(data, "\n")
	for i := 0; i < len(lines); i++ {
		var line = strings.Trim(lines[i], " \t\n\r")
		if len(line) == 0 {
			continue
		}

		// COMMENT / SECTION to be ignore
		if line[0] == '#' || (line[0] == '[' && line[len(line)-1] == ']') {
			continue
		}

		// parse Key
		idx := strings.Index(line, "=")
		if idx < 0 {
			continue
		}
		key := strings.Trim(line[0:idx], " \t\n\r") // trim space\tab\newLine

		// parse Value
		value := strings.Trim(line[idx+1:], " \t\n\r") // trim space\tab\newLine, may be empty string
		if len(value) != 0 {
			// check all value must be use double quotes
			if value[0] != '"' {
				log.Fatalf("value of key `%s` in TOML file `%s` must be using double quotes", key, path)
			}
			pValue, _, ok := p.parseString(value)
			if !ok {
				log.Fatalf("value of key `%s` in TOML file `%s` parse faild, backslash(\\) may be used incorrectly", key, path)
			}
			value = pValue
		}

		// set to map
		if _, exist := p.localesMap[locale]; !exist {
			p.localesMap[locale] = make(map[string]string, 0)
		}

		// check key exist then notice
		if _, exist := p.localesMap[locale][key]; exist {
			log.Printf("Duplicate key-value pairs for key `%s` at file `%s` with locale `%s`", key, path, locale)
		}
		p.localesMap[locale][key] = value
	}
}

// parseString trans value
func (p *Parser) parseString(s string) (string, int, bool) {
	if len(s) <= 0 {
		return "", 0, true // allow empty value
	}

	index := 0
	escape := false
	result := ""
	state := 0 // 0 = left, 1 = inside
	i := -1
	for _, c := range s {
		i++

		if state == 0 {
			if c != '"' {
				return "", 0, false
			}
			state = 1
			continue
		}

		if state == 1 {
			if escape {
				if c == '0' {
					result += "\x00"
				} else if c == 't' {
					result += "\t"
				} else if c == 'n' {
					result += "\n"
				} else if c == 'r' {
					result += "\r"
				} else if c == '"' {
					result += "\""
				} else if c == '\\' {
					result += "\\"
				} else {
					return "", 0, false
				}
				escape = false
				continue
			}

			if c == '\\' {
				escape = true
				continue
			}

			if c == '"' && !escape {
				index = i + 1
				break
			}

			result += string(c)
		}
	}

	return result, index, true
}
