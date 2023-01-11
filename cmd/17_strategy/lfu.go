package strategy

import "fmt"

/*
具象ストラテジー

Least Frequently Used (LFU)： 利用頻度が最低の項目を削除。
*/
type Lfu struct{}

func (l *Lfu) evict(c *Cache) {
	fmt.Println("Evicting by lfu strategy")
}
