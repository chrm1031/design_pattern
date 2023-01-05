package facade

import "fmt"

/*
複雑なサブシステム （部分）
*/
type Notification struct{}

func (n *Notification) sendWalletCreditNotification() {
	fmt.Println("Sending wallet credit notification")
}

func (n *Notification) sendWalletDebitNotification() {
	fmt.Println("Sending wallet debit notification")
}
