package observer

/*
オブザーバー
*/
type Observer interface {
	update(string)
	getID() string
}
