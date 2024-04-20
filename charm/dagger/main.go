// A generated module for Charm functions

package main

import (
	"context"
	"strings"
)

type Charm struct{}

// Executes a charm VHS command that generates a gif from a demo.tape file
func (c *Charm) Vhs(
	ctx context.Context,
	// tape file to generate gif
	fileTape *File,
) *File {
	return c.baseVhs(fileTape).
		WithMountedFile("demo.tape", fileTape).
		WithExec([]string{"demo.tape", "-o", "output.gif"}).
		File("output.gif")
}

// Publish a gif file to the VHS service
func (c *Charm) VhsPublish(
	ctx context.Context,
	file *File,
) (string, error) {
	return c.baseVhs(file).
		WithMountedFile("demo.gif", file).
		WithExec([]string{"publish", "demo.gif"}).
		Stdout(ctx)
}

func (c *Charm) baseVhs(file *File) *Container {
	return dag.Container().
		From("ghcr.io/charmbracelet/vhs")
}

// Executes a charm freeze command that generate images from a source code file
func (c *Charm) Freeze(
	ctx context.Context,
	// Source code file to generate images
	file *File,
	// +optional
	params string,
) *File {
	fileName, _ := file.Name(ctx)

	nameFileExport := "freeze.png"
	paramsToAdd := []string{}

	paramsData := strings.Split(params, " ")
	for index, value := range paramsData {
		if value != "" {
			if value == "-o" || value == "--output" {
				nameFileExport = paramsData[index+1]
			}
			paramsToAdd = append(paramsToAdd, value)
		}
	}

	return dag.Container().
		From("alpine:latest").
		WithMountedFile(fileName, file).
		WithExec([]string{"apk", "update"}).
		WithExec([]string{"apk", "add", "wget"}).
		WithExec([]string{"wget", "https://github.com/charmbracelet/freeze/releases/download/v0.1.6/freeze_0.1.6_Linux_x86_64.tar.gz"}).
		WithExec([]string{"tar", "-xvf", "freeze_0.1.6_Linux_x86_64.tar.gz"}).
		WithExec(append([]string{"freeze_0.1.6_Linux_x86_64/freeze", fileName}, paramsToAdd...)).
		File(nameFileExport)
}
