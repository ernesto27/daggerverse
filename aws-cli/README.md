# Dagger AWS Module


### CLI example

Execute any aws cli command

```shell
dagger call -m github.com/ernesto27/daggerverse/aws-cli run \
 --command="s3 ls"  \
 --dir-config ~/.aws/
```

Created image and publish to ECR

```shell
dagger call -m github.com/ernesto27/daggerverse/aws-cli publish-to-ecr \
 --dir-config ~/.aws \
 --dir-source . \
 --region="us-west-2" \
 --registry="registry-url" \
 --uri="uri"
```

Update ECS service
    
```shell
dagger call -m github.com/ernesto27/daggerverse/aws-cli update-ecs-service \
 --dir-config ~/.aws \
 --region="us-west-2" \
 --task-definition ./task-definition.json \
 --cluster="your-cluster" \
 --service="your-service" \
 --task-definition-name="your-td"
```

### Golang

Install module 

```shell
dagger install github.com/ernesto27/daggerverse/aws-cli
```

```go
package main

import (
	"context"
)

type Example struct{}

func (m *Demo) Example(ctx context.Context, command string, dirConfig *Directory) (string, error) {
	return dag.
			AwsCli().
			Run(ctx, command, dirConfig)
}
```