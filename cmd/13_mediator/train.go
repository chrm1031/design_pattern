package mediator

/*
コンポーネント
*/
type Train interface {
	arrive()
	depart()
	permitArrival()
}
