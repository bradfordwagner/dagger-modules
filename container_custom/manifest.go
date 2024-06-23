package main

import (
	"context"
)

// Manifest - configures manifest in registry. not meant to be run locally
func (m *ContainerCustom) Manifest(
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
		image := imageTag(c, b, version)
		o, err = dag.Lib().ManifestTool(ctx, actor, token, image, b.Architectures)
		if err != nil {
			return
		}
	}

	return
}
