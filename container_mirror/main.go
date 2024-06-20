// A generated module for ContainerMirror functions
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

type ContainerMirror struct{}

type Config struct {
	TargetRepo string `json:"target_repo"`
}

type Build struct {
	Repo          string   `json:"repo"`
	Tag           string   `json:"tag"`
	Architectures []string `json:"archs"`
}

// Returns a container that echoes whatever string argument is provided
// func (m *ContainerMirror) Mirror() *Container {
func (m *ContainerMirror) Mirror(
	ctx context.Context,
	src *Directory,
) string {
	entries, err := src.Entries(ctx)
	if err != nil {
		return err.Error()
	}
	return strings.Join(entries, "\n")
}

// Returns lines that match a pattern in the files of the provided Directory
func (m *ContainerMirror) GrepDir(ctx context.Context, directoryArg *Directory, pattern string) (string, error) {
	return dag.Container().
		From("alpine:latest").
		WithMountedDirectory("/mnt", directoryArg).
		WithWorkdir("/mnt").
		WithExec([]string{"grep", "-R", pattern, "."}).
		Stdout(ctx)
}
