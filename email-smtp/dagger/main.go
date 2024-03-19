// This module send a email using a SMTP server

package main

import (
	"context"
	"strings"

	"gopkg.in/gomail.v2"
)

type EmailSmtp struct{}

// Example usage:
//
//	dagger -m github.com/ernesto27/daggerverse/email-smtp call send  \
//		--from="from@gmail.com" \
//		--to="someemail@gmail.com" \
//		--subject="Hello" \
//		--body="Hello, World!" \
//		--host="smtp.mailtrap.io" \
//		--port=587  \
//		--username env:SMTP_USERNAME \
//		--password env:SMPT_PASSWORD
//		--attachment="path/to/attachment" \
//		--embed="path/to/embed"
func (m *EmailSmtp) Send(
	ctx context.Context,
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
	// +optional
	attachment *File,
	// +optional
	embed *File,
) (string, error) {
	emails := strings.Split(to, ",")

	mail := gomail.NewMessage()
	mail.SetHeader("From", from)
	mail.SetHeader("To", emails...)
	mail.SetHeader("Subject", subject)

	if attachment != nil {
		attachName, err := attachment.Name(ctx)
		if err != nil {
			return "", err
		}
		attachment.Export(ctx, attachName)
		mail.Attach(attachName)
	}

	if embed != nil {
		embedName, err := embed.Name(ctx)
		if err != nil {
			return "", err
		}
		embed.Export(ctx, embedName)
		mail.Embed(embedName)
	}

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
