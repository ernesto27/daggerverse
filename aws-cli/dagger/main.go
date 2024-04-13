// This module executes AWS CLI commands

package main

import (
	"context"
	"strings"
)

type AwsCli struct{}

// Example usage
//
// dagger call -m github.com/ernesto27/daggerverse/aws-cli run
//     --command="s3 ls"  \
//     --dir-config ~/.aws/

func (m *AwsCli) Run(
	// AWS CLI command to execute
	command string,
	// AWS config credentials directory
	dirConfig *Directory,
	// Directory with files to use in the command
	// +optional
	dirFiles *Directory,
) (string, error) {

	commandToExecute := []string{}

	for _, value := range strings.Split(command, " ") {
		if value != "" {
			commandToExecute = append(commandToExecute, value)
		}
	}

	container := dag.Container().
		From("amazon/aws-cli:latest").
		WithMountedDirectory("/root/.aws/", dirConfig)

	if dirFiles != nil {
		container = container.WithMountedDirectory("/aws/", dirFiles)
	}

	resp, err := container.WithExec(commandToExecute).
		Stdout(context.Background())

	return resp, err
}

// Example usage
//
// dagger call push-to-ecr \
// --dir-config ~/.aws \
// --dir-source . \
// --region="us-west-2" \
// --registry="registry-url" \
// --uri="uri"

func (m *AwsCli) PushToECR(
	ctx context.Context,
	// AWS config credentials directory
	dirConfig *Directory,
	// Directory with source code and Dockerfile
	dirSource *Directory,
	// AWS region
	region string,
	// ECR registry URL
	registry string,
	// ECR image URI
	uri string,
) (string, error) {
	token, err := dag.Container().
		From("amazon/aws-cli:latest").
		WithMountedDirectory("/root/.aws/", dirConfig).
		WithExec([]string{"ecr", "get-login-password", "--region", region}).
		Stdout(ctx)
	if err != nil {
		return "", err
	}

	secret := dag.SetSecret("AWS", token)
	return dirSource.DockerBuild().
		WithRegistryAuth(registry, "AWS", secret).
		Publish(ctx, uri)

}
