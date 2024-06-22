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
)

func (m *ContainerMirror) Manifest(
	ctx context.Context,
	src *Directory,
	// +default="latest"
	version string,
	// GitHub actor, --token=env:GITHUB_API_TOKEN,--token=cmd:"gh auth token"
	actor *Secret,
	// GitHub API token, --token=env:GITHUB_API_TOKEN,--token=cmd:"gh auth token"
	token *Secret,
) (o string, err error) {
	c, err := loadConfig(ctx, src)
	if err != nil {
		return
	}

	// run manifest tool for each build
	for _, b := range c.Builds {
		dag.Lib().ManifestTool(ctx, actor, token, image, b.Architectures)
	}

	return
}
