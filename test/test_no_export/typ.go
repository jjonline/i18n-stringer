package test

//go:generate $GOPATH/bin/i18n-stringer -type code_no_export -check
//go:generate $GOPATH/bin/i18n-stringer -type code_no_export -output stringer.go

type code_no_export int

const (
	HELLO code_no_export = iota + 1
	WORLD
)
