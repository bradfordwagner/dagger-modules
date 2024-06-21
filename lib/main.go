// A generated module for Lib functions
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

type Lib struct{}

// Returns lines that match a pattern in the files of the provided Directory
func (m *Lib) OpenConfigYaml(ctx context.Context, src Directory) (bytes []byte, err error) {
	configFile := src.File("config.yaml")
	bytes, err = configFile.Contents(ctx)
	if err != nil {
		return
	}
	return bytes, nil
}
