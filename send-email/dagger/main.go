// This module send a email using a SMTP server

package main

import (
	"context"
	"strings"

	"gopkg.in/gomail.v2"
)

type SendEmail struct{}

// Example usage:
//
//	dagger call send-email  \
//		--from="from@gmail.com" \
//		--to="someemail@gmail.com,otheremail@gmail.com" \
//		--subject="Hello" \
//		--body="Hello, World!" \
//		--host="smtp.mailtrap.io" \
//		--port=587  \
//		--username env:SMTP_USERNAME \
//		--password env:SMPT_PASSWORD
func (m *SendEmail) SendEmail(
	// From email address
	from string,
	// To email address
	to string,
	// Email subject
	subject string,
	// Email body
	body string,
	// SMTP server host
	host string,
	// SMTP server port
	// +optional
	// +default=587
	port int,
	// SMTP username
	username *Secret,
	// SMTP password
	password *Secret,
) (string, error) {

	ctx := context.TODO()
	emails := strings.Split(to, ",")

	mail := gomail.NewMessage()
	mail.SetHeader("From", from)
	mail.SetHeader("To", emails...)
	mail.SetHeader("Subject", subject)

	usernamePlainText, err := username.Plaintext(ctx)
	if err != nil {
		return "", err
	}

	passwordPlainText, err := password.Plaintext(ctx)
	if err != nil {
		return "", err
	}

	d := gomail.NewDialer(host, port, usernamePlainText, passwordPlainText)

	if err := d.DialAndSend(mail); err != nil {
		return "", err
	}

	return "Email sent successfully", nil

}
