package strategy

import "fmt"

/*
具象ストラテジー

First In, First Out (FIFO)： 最初に作られた項目を削除。
*/
type Fifo struct{}

func (l *Fifo) evict(c *Cache) {
	fmt.Println("Evicting by fifo strategy")
}
