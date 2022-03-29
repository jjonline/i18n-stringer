package test_no_export

//go:generate $GOPATH/bin/i18n-stringer -type code_no_export -check
//go:generate $GOPATH/bin/i18n-stringer -type code_no_export -output stringer.go

type code_no_export uint8

const (
	HELLO code_no_export = iota + 1
	WORLD
	ALIAS = WORLD // test for repeat
)
