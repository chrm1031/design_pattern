package observer

/*
サブジェクト
*/
type Subject interface {
	register(Observer)
	deregster(Observer)
	notifyAll()
}
