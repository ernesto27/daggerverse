# Dagger email smtp Module


### CLI example

Send email using SMTP server

```shell
export SMTP_USERNAME=<your-smtp-username>
export SMTP_PASSWORD=<your-smtp-password>

dagger -m github.com/ernesto27/daggerverse/email-smtp \
call send \
--from="from@gmail.com" \
--to="mail@gmail.com" \
--subject="Hello" \
--body="Hello, World!" \
--host="smtp.mailtrap.io" \
--username env:SMTP_USERNAME \ 
--password env:SMTP_PASSWORD \
--attachment="path/to/attachment" \
--embed="path/to/embed"
```


### Golang

Install module 

```shell
dagger install github.com/ernesto27/daggerverse/email-smtp
```

```go
package main

import (
	"context"
)

type Example struct{}

func (m *Demo) Example(ctx context.Context, token *Secret) (string, error) {
	from := "from@gmail.com"
	to := "to@gmail.com"
	subject := "Test Email"
	body := "This is a test email"
	host := "smtp.host"
	return dag.EmailSMTP().Send(context.TODO(), from, to, subject, body, host, username, password)
}
```