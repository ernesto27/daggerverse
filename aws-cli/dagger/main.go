// This module provides a set of functions to interact with AWS CLI

package main

import (
	"context"
	"strings"
)

type AwsCli struct{}

// Executes an AWS CLI command
// Example usage
//
// dagger call -m github.com/ernesto27/daggerverse/aws-cli run \
// --command="s3 ls"  \
// --dir-config ~/.aws/
func (a *AwsCli) Run(
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

	container := a.baseContainer(dirConfig)

	if dirFiles != nil {
		container = container.WithMountedDirectory("/aws/", dirFiles)
	}

	resp, err := container.WithExec(commandToExecute).
		Stdout(context.Background())

	return resp, err
}

// Log in to AWS, build a Docker image and push it to ECR
// Example usage
//
// dagger call -m github.com/ernesto27/daggerverse/aws-cli push-to-ecr \
// --dir-config ~/.aws \
// --dir-source . \
// --region="us-west-2" \
// --registry="registry-url" \
// --uri="uri"
func (a *AwsCli) PublishToEcr(
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
	token, err := a.baseContainer(dirConfig).
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

// Update an ECS service with a new task definition
// Example usage
//
// dagger call -m github.com/ernesto27/daggerverse/aws-cli update-ecs-service \
// --dir-config ~/.aws \
// --region="us-west-2" \
// --task-definition ./task-definition.json \
// --cluster="your-cluster" \
// --service="your-service" \
// --task-definition-name="your-td"
func (a *AwsCli) UpdateEcsService(
	ctx context.Context,
	dirConfig *Directory,
	region string,
	taskDefinition *File,
	cluster string,
	service string,
	taskDefinitionName string,
) (string, error) {
	_, err := a.baseContainer(dirConfig).
		WithMountedFile("task-definition.json", taskDefinition).
		WithExec([]string{"ecs", "register-task-definition", "--region", region, "--cli-input-json", "file://task-definition.json"}).
		Stdout(ctx)

	if err != nil {
		return "", err
	}

	return a.baseContainer(dirConfig).
		WithExec([]string{"ecs", "update-service", "--region", region, "--cluster", cluster, "--service", service, "--task-definition", taskDefinitionName}).
		Stdout(ctx)

}

func (a AwsCli) baseContainer(dirConfig *Directory) *Container {
	return dag.Container().
		From("amazon/aws-cli:latest").
		WithMountedDirectory("/root/.aws/", dirConfig)
}
