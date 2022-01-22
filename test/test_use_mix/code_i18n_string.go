// Code generated by "i18n-stringer -type Code,Test,Single"; DO NOT EDIT.

package test_use_mix

import (
	"context"
	"fmt"
	"strconv"
)

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the i18n-stringer command to generate them again.
	var x [1]struct{}
	_ = x[CodeOK-1]
	_ = x[CodeErr-2]
	_ = x[CodeFail-3]
	_ = x[CodeRange1-20001]
	_ = x[CodeRange2-20002]
	_ = x[CodeRange3-20003]
	_ = x[CodeRange4-20004]
	_ = x[CodeRange5-20204]
	_ = x[CodeRange6-20205]
	_ = x[CodeRange7-20206]
	_ = x[CodeRange8-20206]
	_ = x[CodeRange9-20301]
	_ = x[CodeRange10-20302]
	_ = x[CodeTe1-20310]
	_ = x[CodeTe2-20311]
	_ = x[CodeSe1-20400]
	_ = x[CodeSe2-20401]
	_ = x[CodeSe3-20410]
	_ = x[CodeSe4-20411]
	_ = x[CodeAe1-20420]
	_ = x[CodeAe2-20421]
	_ = x[CodeBe1-20430]
	_ = x[CodeBe2-20431]
	_ = x[CodeCe1-20440]
	_ = x[CodeCe2-20441]
	_ = x[CodeDe1-20450]
	_ = x[CodeDe2-20451]
	_ = x[CodeEe1-20460]
	_ = x[CodeEe2-20461]
	_ = x[CodeFe1-20470]
	_ = x[CodeFe2-20471]
	_ = x[CodeFe3-20472]
	_ = x[CodeGe1-20480]
	_ = x[CodeGe2-20481]
	_ = x[CodeXe1-20490]
	_ = x[CodeXe2-20491]
	_ = x[CodeXe3-20491]
}

const (
	_Code_En_name   = "CodeOKCodeErrCodeFailCodeRange1CodeRange2CodeRange3CodeRange4CodeRange5CodeRange6CodeRange7CodeRange9CodeRange10CodeTe1CodeTe2CodeSe1CodeSe2CodeSe3CodeSe4CodeAe1CodeAe2CodeBe1CodeBe2CodeCe1CodeCe2CodeDe1CodeDe2CodeEe1CodeEe2CodeFe1CodeFe2CodeFe3CodeGe1CodeGe2CodeXe1CodeXe2"
	_Code_ZhHk_name = "CodeOKCodeErrCodeFailCodeRange1CodeRange2CodeRange3CodeRange4CodeRange5CodeRange6CodeRange7CodeRange9CodeRange10CodeTe1CodeTe2CodeSe1CodeSe2CodeSe3CodeSe4CodeAe1CodeAe2CodeBe1CodeBe2CodeCe1CodeCe2CodeDe1CodeDe2CodeEe1CodeEe2CodeFe1CodeFe2CodeFe3CodeGe1CodeGe2CodeXe1CodeXe2"
)

var (
	_Code_En_map = map[Code]string{
		1:     _Code_En_name[0:6],
		2:     _Code_En_name[6:13],
		3:     _Code_En_name[13:21],
		20001: _Code_En_name[21:31],
		20002: _Code_En_name[31:41],
		20003: _Code_En_name[41:51],
		20004: _Code_En_name[51:61],
		20204: _Code_En_name[61:71],
		20205: _Code_En_name[71:81],
		20206: _Code_En_name[81:91],
		20301: _Code_En_name[91:101],
		20302: _Code_En_name[101:112],
		20310: _Code_En_name[112:119],
		20311: _Code_En_name[119:126],
		20400: _Code_En_name[126:133],
		20401: _Code_En_name[133:140],
		20410: _Code_En_name[140:147],
		20411: _Code_En_name[147:154],
		20420: _Code_En_name[154:161],
		20421: _Code_En_name[161:168],
		20430: _Code_En_name[168:175],
		20431: _Code_En_name[175:182],
		20440: _Code_En_name[182:189],
		20441: _Code_En_name[189:196],
		20450: _Code_En_name[196:203],
		20451: _Code_En_name[203:210],
		20460: _Code_En_name[210:217],
		20461: _Code_En_name[217:224],
		20470: _Code_En_name[224:231],
		20471: _Code_En_name[231:238],
		20472: _Code_En_name[238:245],
		20480: _Code_En_name[245:252],
		20481: _Code_En_name[252:259],
		20490: _Code_En_name[259:266],
		20491: _Code_En_name[266:273],
	}
	_Code_ZhHk_map = map[Code]string{
		1:     _Code_ZhHk_name[0:6],
		2:     _Code_ZhHk_name[6:13],
		3:     _Code_ZhHk_name[13:21],
		20001: _Code_ZhHk_name[21:31],
		20002: _Code_ZhHk_name[31:41],
		20003: _Code_ZhHk_name[41:51],
		20004: _Code_ZhHk_name[51:61],
		20204: _Code_ZhHk_name[61:71],
		20205: _Code_ZhHk_name[71:81],
		20206: _Code_ZhHk_name[81:91],
		20301: _Code_ZhHk_name[91:101],
		20302: _Code_ZhHk_name[101:112],
		20310: _Code_ZhHk_name[112:119],
		20311: _Code_ZhHk_name[119:126],
		20400: _Code_ZhHk_name[126:133],
		20401: _Code_ZhHk_name[133:140],
		20410: _Code_ZhHk_name[140:147],
		20411: _Code_ZhHk_name[147:154],
		20420: _Code_ZhHk_name[154:161],
		20421: _Code_ZhHk_name[161:168],
		20430: _Code_ZhHk_name[168:175],
		20431: _Code_ZhHk_name[175:182],
		20440: _Code_ZhHk_name[182:189],
		20441: _Code_ZhHk_name[189:196],
		20450: _Code_ZhHk_name[196:203],
		20451: _Code_ZhHk_name[203:210],
		20460: _Code_ZhHk_name[210:217],
		20461: _Code_ZhHk_name[217:224],
		20470: _Code_ZhHk_name[224:231],
		20471: _Code_ZhHk_name[231:238],
		20472: _Code_ZhHk_name[238:245],
		20480: _Code_ZhHk_name[245:252],
		20481: _Code_ZhHk_name[252:259],
		20490: _Code_ZhHk_name[259:266],
		20491: _Code_ZhHk_name[266:273],
	}
)

// _transOne translate one CONST
func (i Code) _transOne(locale string) string {
	switch locale {
	case "en":
		if str, ok := _Code_En_map[i]; ok {
			return str
		}
		return "Code[" + locale + "](" + strconv.FormatInt(int64(i), 10) + ")"
	case "zh-hk":
		if str, ok := _Code_ZhHk_map[i]; ok {
			return str
		}
		return "Code[" + locale + "](" + strconv.FormatInt(int64(i), 10) + ")"
	default:
		// Normally unreachable, should not happen but be cautious
		return ""
	}
}

// _Code_supported All supported locales record
var _Code_supported = map[string]int{"en": 0, "zh-hk": 1}

// _Code_defaultLocale default locale
// generated pass by i18n-stringer flag -defaultlocale, Don't assign directly
var _Code_defaultLocale = "en"

// _Code_ctxKey Key from context.Context Value get locale
// generated pass by i18n-stringer flag -ctxkey, Don't assign directly
var _Code_ctxKey = "i18nLocale"

// WARNING: You should use Trans, Lang, Wrap, WrapWithContext method instead
//  - You should not use this method in an internationalized language environment, as well as method Error.
//  - Because this method always returns the translation value of the default language.
//  - This method implements the fmt.Stringer interface, so that you can output it directly by package fmt,
//  - If you understand the above mechanism then you can use this method with confidence
func (i Code) String() string {
	return i._trans(_Code_defaultLocale)
}

// WARNING: You should use Trans, Lang, Wrap, WrapWithContext method instead
//  - You should not use this method in an internationalized language environment, as well as method String.
//  - Because this method always returns the translation value of the default language.
//  - This method implements the error interface, so that you can return the value as an error,
//  - If you understand the above mechanism then you can use this method with confidence
func (i Code) Error() string {
	return i._trans(_Code_defaultLocale)
}

// Wrap another error with locale set for i18n TYPE Const
//  - err another error
//  - locale i18n locale name
//  - args optional formatting component
func (i Code) Wrap(err error, locale string, args ...Code) *I18nCodeErrorWrap {
	return &I18nCodeErrorWrap{err: err, origin: i, locale: locale, args: args}
}

// WrapWithContext wrap another error with context.Context set for i18n TYPE Const
//  - ctx context with Value use Key from _Code_ctxKey, which pass by i18n-stringer flag -ctxkey
//  - err another error
//  - args optional formatting component
func (i Code) WrapWithContext(ctx context.Context, err error, args ...Code) *I18nCodeErrorWrap {
	return &I18nCodeErrorWrap{err: err, origin: i, locale: _Code_localeFromCtxWithFallback(ctx), args: args}
}

// I18nCodeErrorWrap type i18n error wrapper
//   WARNING
//   This struct ONLY used to wrap the CONST generated by the i18n-stringer tool,
//   Pass easily obtain internationalized translations through Error, String, Translate
//   WARNING
type I18nCodeErrorWrap struct {
	err    error  // wrap another error
	origin Code   // custom shaping type Val
	locale string // i18n locale set
	args   []Code // formatted output replacement component
}

// Translate get translated string
func (e *I18nCodeErrorWrap) Translate() string {
	return e.origin.Trans(e.locale, e.args...)
}

// String implement fmt.Stringer, get translated string use Translate
func (e *I18nCodeErrorWrap) String() string {
	return e.Translate()
}

// Error struct as error, get translated string use Translate
func (e *I18nCodeErrorWrap) Error() string {
	return e.Translate()
}

// Format string form inside error and TOML define
//  - this method will be formatted wrap error. Only for development and debugging
func (e *I18nCodeErrorWrap) Format() string {
	if e.err == nil {
		return e.Error()
	}
	return fmt.Sprintf("%s (%s)", e.Error(), e.err.Error())
}

// Value get original type value
func (e *I18nCodeErrorWrap) Value() Code {
	return e.origin
}

// Unwrap an error. Get the error inside
func (e *I18nCodeErrorWrap) Unwrap() error {
	return e.err
}

// IsLocaleSupport Check if the specified locale is supported
func (i Code) IsLocaleSupport(locale string) bool {
	return _Code_isLocaleSupport(locale)
}

// Lang get target translate text use context.Context
//  - ctx  context with Value use Key from _Code_ctxKey, which pass by i18n-stringer flag -ctxkey
//  - args Optional placeholder replacement value
func (i Code) Lang(ctx context.Context, args ...Code) string {
	return i._trans(_Code_localeFromCtxWithFallback(ctx), args...)
}

// Trans get target translate text use specified language locale identifier
//  - locale specified language locale identifier, need pass by IsLocaleSupport
//  - args Optional placeholder replacement value
func (i Code) Trans(locale string, args ...Code) string {
	if !_Code_isLocaleSupport(locale) {
		locale = _Code_defaultLocale
	}
	return i._trans(locale, args...)
}

func _Code_isLocaleSupport(locale string) bool {
	_, ok := _Code_supported[locale]
	return ok
}

// _Code_localeFromCtxWithFallback retrieves and returns language locale name from context.
// It returns default locale when _Code_isLocaleSupport is false
func _Code_localeFromCtxWithFallback(ctx context.Context) string {
	if ctx == nil {
		return _Code_defaultLocale
	}
	v := ctx.Value(_Code_ctxKey)
	if v == nil {
		return _Code_defaultLocale
	}
	if vv, ok := v.(string); ok && _Code_isLocaleSupport(vv) {
		return vv
	}
	return _Code_defaultLocale
}

// _trans trustworthy parameters inside method
func (i Code) _trans(locale string, args ...Code) string {
	msg := i._transOne(locale)
	if len(args) > 0 {
		var com []interface{}
		for _, arg := range args {
			com = append(com, arg._transOne(locale))
		}
		return fmt.Sprintf(msg, com...)
	}
	return msg
}

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the i18n-stringer command to generate them again.
	var x [1]struct{}
	_ = x[TestCase01-10]
	_ = x[TestCase02-11]
	_ = x[TestCase03-12]
	_ = x[TestCase04-1001]
	_ = x[TestCase05-1002]
	_ = x[TestCase06-1003]
}

const (
	_Test_En_name_0   = "TestCase01TestCase02TestCase03"
	_Test_ZhHk_name_0 = "TestCase01TestCase02TestCase03"
	_Test_En_name_1   = "TestCase04TestCase05TestCase06"
	_Test_ZhHk_name_1 = "TestCase04TestCase05TestCase06"
)

var (
	_Test_En_index_0   = [...]uint8{0, 10, 20, 30}
	_Test_ZhHk_index_0 = [...]uint8{0, 10, 20, 30}
	_Test_En_index_1   = [...]uint8{0, 10, 20, 30}
	_Test_ZhHk_index_1 = [...]uint8{0, 10, 20, 30}
)

// _transOne translate one CONST
func (i Test) _transOne(locale string) string {
	switch locale {
	case "en":
		switch {
		case 10 <= i && i <= 12:
			i -= 10
			return _Test_En_name_0[_Test_En_index_0[i]:_Test_En_index_0[i+1]]
		case 1001 <= i && i <= 1003:
			i -= 1001
			return _Test_En_name_1[_Test_En_index_1[i]:_Test_En_index_1[i+1]]
		default:
			return "Test[" + locale + "](" + strconv.FormatInt(int64(i), 10) + ")"
		}
	case "zh-hk":
		switch {
		case 10 <= i && i <= 12:
			i -= 10
			return _Test_ZhHk_name_0[_Test_ZhHk_index_0[i]:_Test_ZhHk_index_0[i+1]]
		case 1001 <= i && i <= 1003:
			i -= 1001
			return _Test_ZhHk_name_1[_Test_ZhHk_index_1[i]:_Test_ZhHk_index_1[i+1]]
		default:
			return "Test[" + locale + "](" + strconv.FormatInt(int64(i), 10) + ")"
		}
	default:
		// Normally unreachable, should not happen but be cautious
		return ""
	}
}

// _Test_supported All supported locales record
var _Test_supported = map[string]int{"en": 0, "zh-hk": 1}

// _Test_defaultLocale default locale
// generated pass by i18n-stringer flag -defaultlocale, Don't assign directly
var _Test_defaultLocale = "en"

// _Test_ctxKey Key from context.Context Value get locale
// generated pass by i18n-stringer flag -ctxkey, Don't assign directly
var _Test_ctxKey = "i18nLocale"

// WARNING: You should use Trans, Lang, Wrap, WrapWithContext method instead
//  - You should not use this method in an internationalized language environment, as well as method Error.
//  - Because this method always returns the translation value of the default language.
//  - This method implements the fmt.Stringer interface, so that you can output it directly by package fmt,
//  - If you understand the above mechanism then you can use this method with confidence
func (i Test) String() string {
	return i._trans(_Test_defaultLocale)
}

// WARNING: You should use Trans, Lang, Wrap, WrapWithContext method instead
//  - You should not use this method in an internationalized language environment, as well as method String.
//  - Because this method always returns the translation value of the default language.
//  - This method implements the error interface, so that you can return the value as an error,
//  - If you understand the above mechanism then you can use this method with confidence
func (i Test) Error() string {
	return i._trans(_Test_defaultLocale)
}

// Wrap another error with locale set for i18n TYPE Const
//  - err another error
//  - locale i18n locale name
//  - args optional formatting component
func (i Test) Wrap(err error, locale string, args ...Test) *I18nTestErrorWrap {
	return &I18nTestErrorWrap{err: err, origin: i, locale: locale, args: args}
}

// WrapWithContext wrap another error with context.Context set for i18n TYPE Const
//  - ctx context with Value use Key from _Test_ctxKey, which pass by i18n-stringer flag -ctxkey
//  - err another error
//  - args optional formatting component
func (i Test) WrapWithContext(ctx context.Context, err error, args ...Test) *I18nTestErrorWrap {
	return &I18nTestErrorWrap{err: err, origin: i, locale: _Test_localeFromCtxWithFallback(ctx), args: args}
}

// I18nTestErrorWrap type i18n error wrapper
//   WARNING
//   This struct ONLY used to wrap the CONST generated by the i18n-stringer tool,
//   Pass easily obtain internationalized translations through Error, String, Translate
//   WARNING
type I18nTestErrorWrap struct {
	err    error  // wrap another error
	origin Test   // custom shaping type Val
	locale string // i18n locale set
	args   []Test // formatted output replacement component
}

// Translate get translated string
func (e *I18nTestErrorWrap) Translate() string {
	return e.origin.Trans(e.locale, e.args...)
}

// String implement fmt.Stringer, get translated string use Translate
func (e *I18nTestErrorWrap) String() string {
	return e.Translate()
}

// Error struct as error, get translated string use Translate
func (e *I18nTestErrorWrap) Error() string {
	return e.Translate()
}

// Format string form inside error and TOML define
//  - this method will be formatted wrap error. Only for development and debugging
func (e *I18nTestErrorWrap) Format() string {
	if e.err == nil {
		return e.Error()
	}
	return fmt.Sprintf("%s (%s)", e.Error(), e.err.Error())
}

// Value get original type value
func (e *I18nTestErrorWrap) Value() Test {
	return e.origin
}

// Unwrap an error. Get the error inside
func (e *I18nTestErrorWrap) Unwrap() error {
	return e.err
}

// IsLocaleSupport Check if the specified locale is supported
func (i Test) IsLocaleSupport(locale string) bool {
	return _Test_isLocaleSupport(locale)
}

// Lang get target translate text use context.Context
//  - ctx  context with Value use Key from _Test_ctxKey, which pass by i18n-stringer flag -ctxkey
//  - args Optional placeholder replacement value
func (i Test) Lang(ctx context.Context, args ...Test) string {
	return i._trans(_Test_localeFromCtxWithFallback(ctx), args...)
}

// Trans get target translate text use specified language locale identifier
//  - locale specified language locale identifier, need pass by IsLocaleSupport
//  - args Optional placeholder replacement value
func (i Test) Trans(locale string, args ...Test) string {
	if !_Test_isLocaleSupport(locale) {
		locale = _Test_defaultLocale
	}
	return i._trans(locale, args...)
}

func _Test_isLocaleSupport(locale string) bool {
	_, ok := _Test_supported[locale]
	return ok
}

// _Test_localeFromCtxWithFallback retrieves and returns language locale name from context.
// It returns default locale when _Test_isLocaleSupport is false
func _Test_localeFromCtxWithFallback(ctx context.Context) string {
	if ctx == nil {
		return _Test_defaultLocale
	}
	v := ctx.Value(_Test_ctxKey)
	if v == nil {
		return _Test_defaultLocale
	}
	if vv, ok := v.(string); ok && _Test_isLocaleSupport(vv) {
		return vv
	}
	return _Test_defaultLocale
}

// _trans trustworthy parameters inside method
func (i Test) _trans(locale string, args ...Test) string {
	msg := i._transOne(locale)
	if len(args) > 0 {
		var com []interface{}
		for _, arg := range args {
			com = append(com, arg._transOne(locale))
		}
		return fmt.Sprintf(msg, com...)
	}
	return msg
}

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the i18n-stringer command to generate them again.
	var x [1]struct{}
	_ = x[Sig01-300]
	_ = x[Sig02-301]
	_ = x[Sig03-302]
	_ = x[Sig04-303]
	_ = x[Sig05-304]
	_ = x[Sig06-305]
	_ = x[Sig07-306]
	_ = x[Sig08-307]
	_ = x[Sig09-308]
	_ = x[Sig10-309]
	_ = x[Sig11-310]
	_ = x[Sig12-311]
	_ = x[Sig13-312]
	_ = x[Sig14-313]
	_ = x[Sig15-314]
	_ = x[Sig16-315]
	_ = x[Sig17-316]
	_ = x[Sig18-317]
	_ = x[Sig19-318]
	_ = x[Sig20-319]
	_ = x[Sig21-320]
	_ = x[Sig22-321]
	_ = x[Sig23-322]
	_ = x[Sig24-323]
}

const (
	_Single_En_name   = "Sig01Sig02Sig03Sig04Sig05Sig06Sig07Sig08Sig09Sig10Sig11Sig12Sig13Sig14Sig15Sig16Sig17Sig18Sig19Sig20Sig21Sig22Sig23Sig24"
	_Single_ZhHk_name = "Sig01Sig02Sig03Sig04Sig05Sig06Sig07Sig08Sig09Sig10Sig11Sig12Sig13Sig14Sig15Sig16Sig17Sig18Sig19Sig20Sig21Sig22Sig23Sig24"
)

var (
	_Single_En_index   = [...]uint8{0, 5, 10, 15, 20, 25, 30, 35, 40, 45, 50, 55, 60, 65, 70, 75, 80, 85, 90, 95, 100, 105, 110, 115, 120}
	_Single_ZhHk_index = [...]uint8{0, 5, 10, 15, 20, 25, 30, 35, 40, 45, 50, 55, 60, 65, 70, 75, 80, 85, 90, 95, 100, 105, 110, 115, 120}
)

// _transOne translate one CONST
func (i Single) _transOne(locale string) string {
	i -= 300
	if i < 0 || i >= Single(len(_Single_En_index)-1) {
		return "Single[" + locale + "](" + strconv.FormatInt(int64(i), 10) + ")"
	}

	switch locale {
	case "en":
		return _Single_En_name[_Single_En_index[i]:_Single_En_index[i+1]]
	case "zh-hk":
		return _Single_ZhHk_name[_Single_ZhHk_index[i]:_Single_ZhHk_index[i+1]]
	default:
		// Normally unreachable, should not happen but be cautious
		return ""
	}
}

// _Single_supported All supported locales record
var _Single_supported = map[string]int{"en": 0, "zh-hk": 1}

// _Single_defaultLocale default locale
// generated pass by i18n-stringer flag -defaultlocale, Don't assign directly
var _Single_defaultLocale = "en"

// _Single_ctxKey Key from context.Context Value get locale
// generated pass by i18n-stringer flag -ctxkey, Don't assign directly
var _Single_ctxKey = "i18nLocale"

// WARNING: You should use Trans, Lang, Wrap, WrapWithContext method instead
//  - You should not use this method in an internationalized language environment, as well as method Error.
//  - Because this method always returns the translation value of the default language.
//  - This method implements the fmt.Stringer interface, so that you can output it directly by package fmt,
//  - If you understand the above mechanism then you can use this method with confidence
func (i Single) String() string {
	return i._trans(_Single_defaultLocale)
}

// WARNING: You should use Trans, Lang, Wrap, WrapWithContext method instead
//  - You should not use this method in an internationalized language environment, as well as method String.
//  - Because this method always returns the translation value of the default language.
//  - This method implements the error interface, so that you can return the value as an error,
//  - If you understand the above mechanism then you can use this method with confidence
func (i Single) Error() string {
	return i._trans(_Single_defaultLocale)
}

// Wrap another error with locale set for i18n TYPE Const
//  - err another error
//  - locale i18n locale name
//  - args optional formatting component
func (i Single) Wrap(err error, locale string, args ...Single) *I18nSingleErrorWrap {
	return &I18nSingleErrorWrap{err: err, origin: i, locale: locale, args: args}
}

// WrapWithContext wrap another error with context.Context set for i18n TYPE Const
//  - ctx context with Value use Key from _Single_ctxKey, which pass by i18n-stringer flag -ctxkey
//  - err another error
//  - args optional formatting component
func (i Single) WrapWithContext(ctx context.Context, err error, args ...Single) *I18nSingleErrorWrap {
	return &I18nSingleErrorWrap{err: err, origin: i, locale: _Single_localeFromCtxWithFallback(ctx), args: args}
}

// I18nSingleErrorWrap type i18n error wrapper
//   WARNING
//   This struct ONLY used to wrap the CONST generated by the i18n-stringer tool,
//   Pass easily obtain internationalized translations through Error, String, Translate
//   WARNING
type I18nSingleErrorWrap struct {
	err    error    // wrap another error
	origin Single   // custom shaping type Val
	locale string   // i18n locale set
	args   []Single // formatted output replacement component
}

// Translate get translated string
func (e *I18nSingleErrorWrap) Translate() string {
	return e.origin.Trans(e.locale, e.args...)
}

// String implement fmt.Stringer, get translated string use Translate
func (e *I18nSingleErrorWrap) String() string {
	return e.Translate()
}

// Error struct as error, get translated string use Translate
func (e *I18nSingleErrorWrap) Error() string {
	return e.Translate()
}

// Format string form inside error and TOML define
//  - this method will be formatted wrap error. Only for development and debugging
func (e *I18nSingleErrorWrap) Format() string {
	if e.err == nil {
		return e.Error()
	}
	return fmt.Sprintf("%s (%s)", e.Error(), e.err.Error())
}

// Value get original type value
func (e *I18nSingleErrorWrap) Value() Single {
	return e.origin
}

// Unwrap an error. Get the error inside
func (e *I18nSingleErrorWrap) Unwrap() error {
	return e.err
}

// IsLocaleSupport Check if the specified locale is supported
func (i Single) IsLocaleSupport(locale string) bool {
	return _Single_isLocaleSupport(locale)
}

// Lang get target translate text use context.Context
//  - ctx  context with Value use Key from _Single_ctxKey, which pass by i18n-stringer flag -ctxkey
//  - args Optional placeholder replacement value
func (i Single) Lang(ctx context.Context, args ...Single) string {
	return i._trans(_Single_localeFromCtxWithFallback(ctx), args...)
}

// Trans get target translate text use specified language locale identifier
//  - locale specified language locale identifier, need pass by IsLocaleSupport
//  - args Optional placeholder replacement value
func (i Single) Trans(locale string, args ...Single) string {
	if !_Single_isLocaleSupport(locale) {
		locale = _Single_defaultLocale
	}
	return i._trans(locale, args...)
}

func _Single_isLocaleSupport(locale string) bool {
	_, ok := _Single_supported[locale]
	return ok
}

// _Single_localeFromCtxWithFallback retrieves and returns language locale name from context.
// It returns default locale when _Single_isLocaleSupport is false
func _Single_localeFromCtxWithFallback(ctx context.Context) string {
	if ctx == nil {
		return _Single_defaultLocale
	}
	v := ctx.Value(_Single_ctxKey)
	if v == nil {
		return _Single_defaultLocale
	}
	if vv, ok := v.(string); ok && _Single_isLocaleSupport(vv) {
		return vv
	}
	return _Single_defaultLocale
}

// _trans trustworthy parameters inside method
func (i Single) _trans(locale string, args ...Single) string {
	msg := i._transOne(locale)
	if len(args) > 0 {
		var com []interface{}
		for _, arg := range args {
			com = append(com, arg._transOne(locale))
		}
		return fmt.Sprintf(msg, com...)
	}
	return msg
}
