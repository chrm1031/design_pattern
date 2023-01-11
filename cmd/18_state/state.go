package state

/*
状態インターフェース
*/
type State interface {
	addItem(int) error
	requestItem() error
	insertMoney(int) error
	dispenseItem() error
}
