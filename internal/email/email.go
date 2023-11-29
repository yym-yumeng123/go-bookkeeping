package email

import "gopkg.in/gomail.v2"

func Send() {
	m := gomail.NewMessage()
	m.SetHeader("From", "1614527443@qq.com")
	m.SetHeader("To", "18026493873@163.com")
	// m.SetAddressHeader("Cc", "dan@example.com", "Dan")
	m.SetHeader("Subject", "来自GO")
	m.SetBody("text/html", "Hello <b>杨雨蒙</b>")

	d := gomail.NewDialer("smtp.qq.com", 587, "1614527443@qq.com", "bjgmkibtdxdvcfhg")

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}
