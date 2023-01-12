package command

/*
具象コマンド
*/
type OnCommand struct {
	device Device
}

func (c *OnCommand) execute() {
	c.device.on()
}
