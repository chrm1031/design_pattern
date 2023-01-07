package decorator

/*
具象コンポーネント
*/
type VeggeMania struct{}

func (p *VeggeMania) getPrice() int {
	return 15
}
