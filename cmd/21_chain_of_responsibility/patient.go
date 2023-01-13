package chainofresponsibility

/*
クライアント
*/
// 患者
type Patient struct {
	name              string
	registrationDone  bool
	doctorCheckUpDone bool
	medicineDone      bool
	paymentDone       bool
}
