package strategy

import "fmt"

/*
具象ストラテジー

Least Recently Used (LRU)： 使用後最も時間の経った項目を削除。
*/
type Lru struct{}

func (l *Lru) evict(c *Cache) {
	fmt.Println("Evicting by lru strategy")
}
