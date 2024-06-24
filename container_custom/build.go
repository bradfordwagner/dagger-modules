package main

import (
	"context"
	"dagger/container-custom/internal/dagger"
	"encoding/json"
	"fmt"
	"strings"
)

// Build - builds the container image
func (m *ContainerCustom) Build(
	ctx context.Context,
	// +default=0
	index int,
	src *Directory,
	// +default="latest"
	version string,
	// +default=true
	isDev bool,
) (o string, err error) {
	// generate products
	products, err := m.Product(ctx, src, version)
	if err != nil {
		return
	}

	// load product config
	product := products[index]
	b, err := json.Marshal(product)
	if err != nil {
		return
	}
	productJson := string(b)

	// set target image
	target, err := dag.Lib().ArchImageName(ctx, product.TargetImage, product.Architecture)
	if err != nil {
		return
	}

	//dockerfile setup
	dockerfile := fmt.Sprintf(`
			FROM %s
			`, product.UpstreamImage)
	d := src.WithNewFile("Dockerfile", dockerfile)
	container := d.DockerBuild(dagger.DirectoryDockerBuildOpts{
		Platform: dagger.Platform(product.Architecture),
	})
	o = strings.Join([]string{target, productJson, dockerfile}, "\n")

	// publish only through pipeline
	if !isDev {
		_, err = container.Publish(ctx, target)
	}

	return
}
