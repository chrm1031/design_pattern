package templatemethod

/*
テンプレート・メソッド
*/
type IOtp interface {
	genRandomOTP(int) string       // n桁の乱数を発生
	saveOTPCache(string)           // のちの検証のためキャッシュを保存
	getMessage(string) string      // 通知内容を用意
	sendNotification(string) error // 通知を送信
}

type Otp struct {
	iOtp IOtp
}

func (o *Otp) genAndSendOTP(otpLength int) error {
	otp := o.iOtp.genRandomOTP(otpLength)
	o.iOtp.saveOTPCache(otp)
	message := o.iOtp.getMessage(otp)
	if err := o.iOtp.sendNotification(message); err != nil {
		return err
	}
	return nil
}
