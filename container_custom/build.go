package main

import (
	"context"
	"dagger/container-custom/internal/dagger"
	"encoding/json"
	"strings"
)

// Build - builds the container image
func (m *ContainerCustom) Build(
	ctx context.Context,
	// +default=0
	index int,
	src *dagger.Directory,
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

	// find build to pull args
	c, err := loadConfig(ctx, src)
	if err != nil {
		return
	}
	var build Build
	for _, b := range c.Builds {
		if b.OS == product.OS {
			build = b
			break
		}
	}

	//dockerfile setup
	var buildArgs []dagger.BuildArg
	for k, v := range build.Args {
		buildArgs = append(buildArgs, dagger.BuildArg{Name: k, Value: v})
	}
	buildArgs = append(buildArgs, dagger.BuildArg{Name: "OS", Value: product.UpstreamImage})

	container, err := src.DockerBuild(dagger.DirectoryDockerBuildOpts{
		Platform:  dagger.Platform(product.Architecture),
		BuildArgs: buildArgs,
	}).WithFocus().Sync(ctx)

	o = strings.Join([]string{target, productJson}, "\n")

	// publish only through pipeline
	if !isDev {
		_, err = container.Publish(ctx, target)
	}

	return
}
