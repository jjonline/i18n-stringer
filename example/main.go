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
	fmt.Println("testCase-1")
	fmt.Println(lang.MerchantLoginInvalid.Trans("zh_cn", lang.ComUserName))
	fmt.Println(lang.Ok.Trans("zh_cn"))

	// Specify the language locale invalid, default language translation will be output
	fmt.Println("testCase-2")
	fmt.Println(lang.MerchantLoginInvalid.Trans("zh-cn", lang.ComUserName))

	// can not be formatted
	fmt.Println("testCase-3")
	fmt.Println(lang.MerchantLoginInvalid.String())

	// call String method
	fmt.Println("testCase-4")
	fmt.Println(lang.SubDirectoryTest)

	// Specify the language locale use context.Context param
	fmt.Println("testCase-5")
	ctx := context.WithValue(context.TODO(), "i18n", "en") // Attention the VALUE of second parameter
	fmt.Println(lang.MerchantLoginInvalid.Lang(ctx, lang.ComUserName))
	fmt.Println(lang.Ok.Lang(ctx))
}
