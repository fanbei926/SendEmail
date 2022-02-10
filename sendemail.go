package sendemail

import (
	"crypto/tls"

	"gopkg.in/gomail.v2"
)

/**
* Send an text/plain email
**/
func SendMail(title, host, user, passwd, toAddress, content string, port int) (err error) {
	mail := gomail.NewMessage()
	mail.SetHeader("From", mail.FormatAddress(user, user))
	mail.SetHeader("To", mail.FormatAddress(toAddress, ""))

	mail.SetHeader("Subject", title)
	mail.SetBody("text/plain", content)

	conn := gomail.NewDialer(host, port, user, passwd)
	conn.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	// conn := gomail.Dialer{
	// 	Host:      conf.MAIL.HOST,
	// 	Port:      conf.MAIL.PORT,
	// 	Username:  conf.MAIL.USER,
	// 	Password:  conf.MAIL.PASSWD,
	// 	SSL:       conf.MAIL.PORT == 465,
	// 	TLSConfig: &tls.Config{ServerName: conf.MAIL.HOST, InsecureSkipVerify: true},
	// }
	err = conn.DialAndSend(mail)
	return err
}
