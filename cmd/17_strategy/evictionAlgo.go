package strategy

/*
ストラテジー・インターフェース
*/
type EvictionAlgo interface {
	evict(*Cache)
}
