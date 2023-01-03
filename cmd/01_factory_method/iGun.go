package factorymethod

/*
プロダクト・インターフェース
*/
type IGun interface {
	setName(string)
	setPower(int)
	getName() string
	getPower() int
}
