package email

import (
	"fmt"

	"gopkg.in/gomail.v2"
)

func newDialer() *gomail.Dialer {
	return gomail.NewDialer("smtp.qq.com", 587, "1614527443@qq.com", "bjgmkibtdxdvcfhg")
}

func NewMessage(to, subject, body string) *gomail.Message {
	m := gomail.NewMessage()
	m.SetHeader("From", "1614527443@qq.com")
	m.SetHeader("To", to)
	// m.SetAddressHeader("Cc", "dan@example.com", "Dan")
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)
	return m
}

func Send() {
	m := NewMessage("1614527443@qq.com", "我是默认", "Hello <b>杨雨蒙</b>")
	d := newDialer()

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}

func SendValidationCode(email, code string) error {
	m := NewMessage(
		email,
		fmt.Sprintf("[%s] 记账验证码", code),
		fmt.Sprintf("你正在登录记账网站, 你的验证码是 [%s]", code),
	)
	d := newDialer()
	// Send the email to Bob, Cora and Dan.
	return d.DialAndSend(m)
}
