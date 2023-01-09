package iterator

/*
イテレーター
*/
type Iterator interface {
	hasNext() bool
	getNext() *User
}
