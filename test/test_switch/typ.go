package test_switch

//go:generate $GOPATH/bin/i18n-stringer -type RuneOne,RuneMulti,RuneMap -check
//go:generate $GOPATH/bin/i18n-stringer -type RuneOne,RuneMulti,RuneMap

type RuneOne int
type RuneMulti int
type RuneMap int

const (
	CostRuneOT1 RuneOne = iota + 20
	CostRuneOT2
	CostRuneOT3
)

const (
	CostRuneMT1 RuneMulti = iota + 1
	CostRuneMT2
	CostRuneMT3
	CostRuneMT4 RuneMulti = 1000 + iota
	CostRuneMT5
)

const (
	ConstRuneMaT1  RuneMap = 1000
	ConstRuneMaT2  RuneMap = 2000
	ConstRuneMaT3  RuneMap = 3000
	ConstRuneMaT4  RuneMap = 4000
	ConstRuneMaT5  RuneMap = 5000
	ConstRuneMaT6  RuneMap = 6000
	ConstRuneMaT7  RuneMap = 7000
	ConstRuneMaT8  RuneMap = 8000
	ConstRuneMaT9  RuneMap = 9000
	ConstRuneMaT10 RuneMap = 10000
	ConstRuneMaT11 RuneMap = 11000
)
