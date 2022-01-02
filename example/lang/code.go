package lang

// First check
//go:generate $GOPATH/bin/i18n-stringer -type ErrorCode -tomlpath i18n -check

// Second generation
//go:generate $GOPATH/bin/i18n-stringer -type ErrorCode -defaultlocale zh_cn -ctxkey i18n -tomlpath i18n -defaultlocale en -output i118n_stringer.go

// ErrorCode Define a custom shaping type that supports i18n as an error code
type ErrorCode uint

// Specify the error code of a part of the range as the replacement value content of the formatted output
// NOTE: The range 10000 to 20000 is used as a formatted output replacement component
const (
	ComUserName ErrorCode = 10000 + iota
	ComUserPwd
)

// Assign error code value range
const (
	Ok ErrorCode = 20000 + iota
	UserLoginInvalid
	MerchantLoginInvalid
	SubDirectoryTest
)
