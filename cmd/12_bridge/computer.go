package bridge

/*
抽象化
*/
type Computer interface {
	Print()
	SetPrinter(Printer)
}
