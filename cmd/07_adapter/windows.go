package adapter

import "fmt"

/*
不明なサービス
*/
type Windows struct{}

func (w *Windows) insertIntoUSBPort() {
	fmt.Println("USB connector is plugged into windows machine.")
}
