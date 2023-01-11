package strategy

/*
クライアント・コード
*/
func main() {
	lfu := &Lfu{}
	cache := initCache(lfu)

	cache.add("a", "1")
	cache.add("b", "2")
	cache.add("c", "3")

	lru := &Lru{}
	cache.setEvictionAlgo(lru)
	cache.add("d", "4")

	fifo := &Fifo{}
	cache.setEvictionAlgo(fifo)

	cache.add("e", "5")
}

/*
実行結果
Evicting by lfu strategy
Evicting by lru strategy
Evicting by fifo strategy
*/
