package bridge

import "fmt"

/*
具象実装
*/
type Epson struct{}

func (p *Epson) PrintFile() {
	fmt.Println("Printing by a EPSON Printer")
}
