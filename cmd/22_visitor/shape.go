package visitor

/*
要素
*/
type Shape interface {
	getType() string
	accept(Visitor)
}