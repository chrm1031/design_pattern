package command

/*
受け手インターフェース
*/
type Device interface {
	on()
	off()
}
