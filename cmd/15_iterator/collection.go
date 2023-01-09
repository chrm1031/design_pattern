package iterator

/*
コレクション
*/
type Collection interface {
	createIterator() Iterator
}
