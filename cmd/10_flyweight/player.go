package flyweight

/*
コンテキスト
*/
type Player struct {
	dress      Dress
	playerType string // T（テロリスト） か CT（反テロリスト）
	lat        int
	long       int
}

func newPlayer(playerType, dressType string) *Player {
	dress, _ := getDressFactorySingleInstance().getDressByType(dressType)
	return &Player{
		playerType: playerType,
		dress:      dress,
	}
}

func (p *Player) newLocation(lat, long int) {
	p.lat = lat
	p.long = long
}
