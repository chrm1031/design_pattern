package state

import (
	"fmt"
	"log"
)

/*
クライアント・コード
*/
func main() {
	vendingMachine := newVendingMachine(1, 10)
	if err := vendingMachine.requestItem(); err != nil {
		log.Fatalf(err.Error())
	}
	if err := vendingMachine.insertMoney(10); err != nil {
		log.Fatalf(err.Error())
	}
	if err := vendingMachine.dispenseItem(); err != nil {
		log.Fatalf(err.Error())
	}

	fmt.Println()

	if err := vendingMachine.addItem(2); err != nil {
		log.Fatalf(err.Error())
	}

	fmt.Println()

	if err := vendingMachine.requestItem(); err != nil {
		log.Fatalf(err.Error())
	}
	if err := vendingMachine.insertMoney(10); err != nil {
		log.Fatalf(err.Error())
	}
	if err := vendingMachine.dispenseItem(); err != nil {
		log.Fatalf(err.Error())
	}
}

/*
実行結果
Item requestd
Money entered is ok
Dispensing Item

Adding 2 items

Item requestd
Money entered is ok
Dispensing Item
*/
