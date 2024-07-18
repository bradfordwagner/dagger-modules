package main

import (
	"context"
	"dagger/container-mirror/internal/dagger"
)

// Manifest - configures manifest in registry. not meant to be run locally
func (m *ContainerMirror) Manifest(
	ctx context.Context,
	src *dagger.Directory,
	// +default="latest"
	version string,
	// GitHub actor, --token=env:GITHUB_API_TOKEN,--token=cmd:"gh auth token"
	actor *dagger.Secret,
	// GitHub API token, --token=env:GITHUB_API_TOKEN,--token=cmd:"gh auth token"
	token *dagger.Secret,
) (o string, err error) {
	c, err := loadConfig(ctx, src)
	if err != nil {
		return
	}

	// run manifest tool for each build
	for _, b := range c.Builds {
		image := imageTag(c, b, version)
		o, err = dag.Lib().ManifestTool(ctx, actor, token, image, b.Architectures)
		if err != nil {
			return
		}
	}

	return
}
