package command

/*
具象コマンド
*/
type OffCommand struct {
	device Device
}

func (c *OffCommand) execute() {
	c.device.off()
}
