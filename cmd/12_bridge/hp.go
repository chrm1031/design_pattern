package bridge

import "fmt"

/*
具象実装
*/
type Hp struct{}

func (p *Hp) PrintFile() {
	fmt.Println("Printing by a HP Printer")
}
