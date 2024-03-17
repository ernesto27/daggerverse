# Dagger email smtp Module


### CLI example

Send email using SMTP server

```shell
export SMTP_USERNAME=<your-smtp-username>
export SMPT_PASSWORD=<your-smtp-password>

dagger -m github.com/ernesto27/daggerverse/email-smtp \
call send \
--from="from@gmail.com" \
--to="mail@gmail.com" \
--subject="Hello" \
--body="Hello, World!" \
--host="smtp.mailtrap.io" \
--username env:SMTP_USERNAME \ 
--password env:SMPT_PASSWORD 
```


### Golang

```go
package main

import (
	"context"
)

type Example struct{}

func (m *Demo) Example(ctx context.Context, token *Secret) error {
	
}
```