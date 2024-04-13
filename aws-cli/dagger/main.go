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
	dirFiles *Directory) (string, error) {

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
