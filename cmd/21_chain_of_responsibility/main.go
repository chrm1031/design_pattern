package chainofresponsibility

/*
クライアント・コード
*/
func main() {
	cachier := &Cashier{}

	medical := &Medical{}
	medical.setNext(cachier)

	doctor := &Doctor{}
	doctor.setNext(medical)

	reception := &Reception{}
	reception.setNext(doctor)

	patient := &Patient{}
	reception.execute(patient)
}

/*
実行結果
Reception registering patient
Doctor checking patient
Medical giving medicine to patient
Cashier getting money from patient patient
*/
