// This module send a email using gomail package

package main

import (
	"gopkg.in/gomail.v2"
)

type SendEmail struct{}

func (m *SendEmail) SendEmail(
	from string,
	to string,
	subject string,
	body string,
	host string,
	port int,
	username string,
	password string,

) (string, error) {

	mail := gomail.NewMessage()
	mail.SetHeader("From", from)
	mail.SetHeader("To", to)
	mail.SetHeader("Subject", subject)
	mail.SetBody("text/html", body)

	d := gomail.NewDialer(host, port, username, password)

	if err := d.DialAndSend(mail); err != nil {
		return "", err
	}

	return "Email sent successfully", nil

}
