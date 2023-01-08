package mediator

/*
メディエーター・インターフェース
*/
type Mediator interface {
	canArrive(Train) bool
	notifyAboutDeparture()
}
