package chainofresponsibility

/*
ハンドラー・インターフェース
*/
// 部門
type Department interface {
	execute(*Patient)
	setNext(Department)
}
