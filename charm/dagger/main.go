// A generated module for Charm functions
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
)

type Charm struct{}

func (c *Charm) Vhs(ctx context.Context, fileTape *File) *File {
	return c.baseVhs(fileTape).
		WithMountedFile("demo.tape", fileTape).
		WithExec([]string{"demo.tape", "-o", "output.gif"}).
		File("output.gif")
}

func (c *Charm) VhsPublish(ctx context.Context, file *File) (string, error) {
	return c.baseVhs(file).
		WithMountedFile("demo.gif", file).
		WithExec([]string{"publish", "demo.gif"}).
		Stdout(ctx)
}

func (c *Charm) baseVhs(file *File) *Container {
	return dag.Container().
		From("ghcr.io/charmbracelet/vhs")
}

func (c *Charm) Freeze(ctx context.Context) (string, error) {
	return dag.Container().
		From("alpine:latest").
		WithExec([]string{"apk", "update"}).
		WithExec([]string{"apk", "add", "wget"}).
		WithExec([]string{"wget", "https://github.com/charmbracelet/freeze/releases/download/v0.1.6/freeze_0.1.6_Linux_x86_64.tar.gz"}).
		WithExec([]string{"mkdir", "freeze"}).
		// WithExec([]string{"tar", "-xvf", "freeze_0.1.6_Linux_x86_64.tar.gz", "--directory", "freeze"}).
		WithExec([]string{"tar", "-xvf", "freeze_0.1.6_Linux_x86_64.tar.gz"}).
		WithExec([]string{"freeze_0.1.6_Linux_x86_64/freeze"}).
		Stdout(ctx)
}
