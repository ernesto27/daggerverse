// A generated module for AwsCli functions
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
