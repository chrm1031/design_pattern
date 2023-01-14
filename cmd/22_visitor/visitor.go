package visitor

/*
ビジター
*/
type Visitor interface {
	visitForSquare(*Square)
	visitForCircle(*Circle)
	visitForRectangle(*Rectangle)
}
