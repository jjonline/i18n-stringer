package test_check_const

// just check const define in TOML files
//go:generate $GOPATH/bin/i18n-stringer -type Code,Test,Single -check

type Code int
type Test int
type Single int

const (
	CodeOK Code = iota + 1
	CodeErr
	CodeFail
	CodeRange1  Code = 20001
	CodeRange2  Code = 20002
	CodeRange3  Code = 20003
	CodeRange4  Code = 20004
	CodeRange5  Code = 20204
	CodeRange6  Code = 20205
	CodeRange7  Code = 20206
	CodeRange8  Code = 20206
	CodeRange9  Code = 20301
	CodeRange10 Code = 20302
	CodeTe1     Code = 20310
	CodeTe2     Code = 20311
	CodeSe1     Code = 20400
	CodeSe2     Code = 20401
	CodeSe3     Code = 20410
	CodeSe4     Code = 20411
	CodeAe1     Code = 20420
	CodeAe2     Code = 20421
	CodeBe1     Code = 20430
	CodeBe2     Code = 20431
	CodeCe1     Code = 20440
	CodeCe2     Code = 20441
	CodeDe1     Code = 20450
	CodeDe2     Code = 20451
	CodeEe1     Code = 20460
	CodeEe2     Code = 20461
	CodeFe1     Code = 20470
	CodeFe2     Code = 20471
	CodeFe3     Code = 20472
	CodeGe1     Code = 20480
	CodeGe2     Code = 20481
	CodeXe1     Code = 20490
	CodeXe2     Code = 20491
	CodeXe3     Code = 20491
)

const (
	TestCase01 Test = iota + 10
	TestCase02
	TestCase03
	TestCase04 Test = 1001
	TestCase05 Test = 1002
	TestCase06 Test = 1003
)

const (
	Sig01 Single = iota + 300
	Sig02
	Sig03
	Sig04
	Sig05
	Sig06
	Sig07
	Sig08
	Sig09
	Sig10
	Sig11
	Sig12
	Sig13
	Sig14
	Sig15
	Sig16
	Sig17
	Sig18
	Sig19
	Sig20
	Sig21
	Sig22
	Sig23
	Sig24
)
