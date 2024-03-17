// A generated module for SendEmail functions
//
// This module has been generated via dagger init and serves as a reference to
// basic module structure as you get started with Dagger.
//
// Two functions have been pre-created. You can modify, delete, or add to them,
// as needed. They demonstrate usage of arguments and return types using simple
// echo and grep commands. The functions can be called from the dagger CLI or
// from one of the SDKs.
//
// The first line in this comment block is a short description line and the
// rest is a long description with more detail on the module's purpose or usage,
// if appropriate. All modules should have a short description.

package main

import (
	"fmt"

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

) string {

	mail := gomail.NewMessage()
	mail.SetHeader("From", from)
	mail.SetHeader("To", to)
	mail.SetHeader("Subject", subject)
	mail.SetBody("text/html", body)

	d := gomail.NewDialer(host, port, username, password)

	if err := d.DialAndSend(mail); err != nil {
		fmt.Println(err)
		return err.Error()
	}

	return ""

}
