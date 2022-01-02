package main

import (
	"context"
	"fmt"
	"github.com/tvb-sz/i18n-stringer/example/lang"
)

func main() {
	// ++++++++++++++++++++
	// NOTE: First you should execute
	//   go generate ./...
	// ++++++++++++++++++++

	// Specify the language locale use string param
	fmt.Println("-----\ntestCase-1")
	fmt.Println(lang.MerchantLoginInvalid.Trans("zh_cn", lang.ComUserName))
	fmt.Println(lang.Ok.Trans("zh_cn"))

	// Specify the language locale invalid, default language translation will be output
	fmt.Println("-----\ntestCase-2")
	fmt.Println(lang.MerchantLoginInvalid.Trans("zh-cn", lang.ComUserName))

	// can not be formatted
	fmt.Println("-----\ntestCase-3")
	fmt.Println(lang.MerchantLoginInvalid.String())

	// call String method
	fmt.Println("-----\ntestCase-4")
	fmt.Println(lang.SubDirectoryTest)

	// Specify the language locale use context.Context param
	fmt.Println("-----\ntestCase-5")
	ctx := context.WithValue(context.TODO(), "i18n", "en") // Attention the VALUE of second parameter
	fmt.Println(lang.MerchantLoginInvalid.Lang(ctx, lang.ComUserName))
	fmt.Println(lang.Ok.Lang(ctx))

	// use ErrorWrap struct
	// JUST used when you need to wrap another error return
	fmt.Println("-----\ntestCase-6")
	wrapErr := fmt.Errorf("this is an error need to be wrapped")
	err := lang.MerchantLoginInvalid.Wrap(wrapErr, "zh_cn")
	fmt.Println(err.Translate()) // can not be formatted
	fmt.Println(err)             // use String method as fmt.Stringer, can not be formatted
	fmt.Println(err.Unwrap())    // get wrapped error

	err1 := lang.MerchantLoginInvalid.Wrap(wrapErr, "zh_cn", lang.ComUserName)
	fmt.Println(err1.Translate()) // can be formatted

	err2 := lang.MerchantLoginInvalid.WrapWithContext(ctx, wrapErr, lang.ComUserName)
	fmt.Println(err2.Translate()) // wrap with context.Context, can be formatted
}
