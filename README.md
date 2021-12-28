# i18n-stringer

i18n-stringer is a golang international language pack tool, refer to the official [stinger](https://github.com/golang/tools/tree/master/cmd/stringer) tool

golang国际化语言包工具，参考了官方stringer工具

> `test`目录给出了多种使用示例，请参考

> The `test` directory gives some usage examples

在項目根目錄下使用`make`命令調試`test`目錄裡的測試用例

use command `make` for examples in the `test` directory

````
make debug
````

# 一、Usage/使用

## 1.1、Build/下载编译

go Version below 1.16/go1.16以下版本
````
go get -u -d github.com/tvb-sz/i18n-stringer
````

go Version 1.16 and above/go1.16及其以上版本
````
go install github.com/tvb-sz/i18n-stringer@latest
````

## 1.2、Define INT Constant/定义INT型常量

````
type Code int

const (
    ERROROFYOU Code = iota + 1   
    ERROROFMINE Code = 10000
)
````

## 1.3、Define Language Package/定义语言包

> Use only TOML format files

* 定义语言包目录：语言包目录位于定义常量源码文件的同级目录下的子目录，默认语言包目录名称为`i18n`
* The language package directory is located in a subdirectory of the same level directory that defines the constant source code file. The default language package directory name is `i18n`
* 语言包目录下使用TOML文件定义i18n翻译的键值对；
* Use TOML files in the language package directory to define key-value pairs for i18n translation
* 语言包目录下若使用子目录，则子目录名将被作为语言类型标记，子目录下TOML文件名和文件数目不做限制，例如：`en`；
* If a subdirectory is used in the language package directory, the name of the subdirectory will be marked as the language type, TOML file name and number of files in subdirectories are not limited. For example: `en`;
* 语言包目录下不使用子目录直接定义TOML文件的，则TOML文件的文件名将被作为语言类型标记，例如：`en.toml`；
* If the TOML file is directly defined in language package directory, the file name of the TOML file will be marked as the language type. For example: `en.toml`;
* 语言包键值对的键名使用常量字面量，上述例子中`ERROROFYOU`就将作为键名；
* Use constant literals for the key names of language pack key-value pairs. In the above example, `ERROROFYOU` will be used as the key name

## 1.4、代码生成/Generate Code

Go to the directory that defines the constant under the terminal and execute it

终端下进入到定义常量的目录直接执行：
````
$GOPATH/bin/i18n-stringer -type Code -tomlpath i18n
````

You can also use the `go generate` command directly in the constant source code

你也可以配合`go generate`指令直接在定于常量的源码里使用
````
// write this code in your golang source code, then use `go generate` command
//go:generate $GOPATH/bin/i18n-stringer -type Code -tomlpath i18n
````

## 1.5、開發調試/Dev and debugging

可以使用`-check`指令參數核對語言包缺失的鍵值對

You can use the `-check` command
to check the missing key-value pairs of the language pack

````
$GOPATH/bin/i18n-stringer -type Code -tomlpath i18n -check
````

带`-check`的命令執行後如果有缺失鍵值對，則可能會输出如下提示以協助開發

After the command with `-check` is executed,
if there are missing key-value pairs,
the following prompt may be output to assist development

````
i18n-stringer: Check Fail
i18n-stringer: The missing key-value pair information as follows
************TYPE for `Code` of locale `zh-hk`************
CodeOK
CodeErr
CodeFail
CodeRange1
CodeRange2
````

## 1.6、指令詳情/command details

Get more help information about commands

获取更多命令使用帮助信息：
````
$GOPATH/bin/i18n-stringer --help
````
Flags
````
Flags:
  -check
        Check for missing constant literals in YAML files by language
  -ctxkey string
        key used by context.Value for get locale; default i18nLocale
  -defaultlocale string
        set default locale name; default naturally sorted first
  -output string
        output file name; default srcdir/<type>_i18n_string.go
  -tags string
        comma-separated list of build tags to apply
  -tomlpath string
        set toml i18n file path; default srcdir/i18n
  -type string
        comma-separated list of type names; must be set

````

> If your GOBIN directory has been added to the environment variable, the above `$GOPATH/bin/` can also be omitted

> 如果你的GOBIN目录已加入环境变量，上述`$GOPATH/bin/`也是可以省略的

## 1.7、調用/Code call

整形自定義類型除了添加`func (i typ) String() string`方法以實現`fmt.Stringer`接口外，
还会會被添加`func (i typ) Trans(locale string, args ...interface{}) string`和`func (i typ) Lang(ctx context.Context, args ...interface{}) string`方法。
使用`Trans`方法指定字符串形式的語言類型例如`en`即可獲取翻譯文本；
也可以使用`context.Context`攜帶語言類型值的上下文作為參數的方法`Lang`獲取翻譯文本，
需要說明的是`context.Context`携带语言类型的键名由`-ctxkey`指定，默認鍵名為`i18nLocale`。

In addition to adding the `func (i typ) String() string` method to implement the `fmt.Stringer` interface,
the shaping custom type will also be
added with `func (i typ) Trans(locale string, args ...interface{}) string` and `func ( i typ) Lang(ctx context.Context, args ...interface{}) string` method. 
Use the `Trans` method to specify the language type in the form of a string, 
such as `en` to get the translated text; 
you can also use the `context.Context` method that carries the
context of the language type value as a parameter to get the translated text.
The key name of `context.Context` carrying language type is specified by `-ctxkey`, 
and the default key name is `i18nLocale`.

因部分翻譯文本中可能會使用諸如`%s`類型的替換佔位符在代碼中實時更改，建議規劃好整形數值區間，
某些區間的值專門用於替換`%s`的。

Because some translation texts may use replacement placeholders such as `%s` to change in the code in real time, 
it is recommended to plan the integer value range, and this range values are specifically used to replace `%s`.

# 二、TOML规范支持/TOML Specification Support

TOML Link : [https://toml.io/en/](https://toml.io/en/)

* TOML文件仅支持`Basic strings`形式的字符串键值对，形如：`Key="value"`，也就意味着一对键值对只能位于一行
* TOML file only support `Basic strings` key/value, shaped like `Key="value"`, do not support `Multi-line basic strings`, pair K/V only be located on one line
* TOML键名仅支持`裸键`，键名只能包含ASCII字母，ASCII数字，下划线和短横线（`A-Za-z0-9_-`）
* TOML file only support `Bare keys`,only contain ASCII letters, ASCII digits, underscores, and dashes(`A-Za-z0-9_-`)
* 区块也就是TOML官方的`Table`将被忽略
* The block section, which is the TOML official `Table`, will be ignored
* 支持`#`开头的注释，注释将被忽略
* Support comments starting with `#`, comments will be ignored
