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
	"dagger/container-mirror/internal/dagger"
	"encoding/json"
	"fmt"
	"strings"
)

func (m *ContainerMirror) Build(
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
	products, err := m.Product(ctx, src)
	if err != nil {
		return
	}

	// load product config
	product := products[index]
	b, err := json.Marshal(product)
	if err != nil {
		return
	}
	productJson = string(b)

	// load config
	c, err := loadConfig(ctx, src)
	if err != nil {
		return
	}
	target := fmt.Sprintf("%s:%s-%s_%s", c.TargetRepo, version, product.Repo, product.Tag)

	//dockerfile setup
	dockerfile := fmt.Sprintf(`
			FROM %s:%s
			`, product.Repo, product.Tag)
	d := src.WithNewFile("Dockerfile", dockerfile)
	container := d.DockerBuild(dagger.DirectoryDockerBuildOpts{
		Platform: dagger.Platform(product.Architecture),
	})
	o = strings.Join([]string{productJson, dockerfile}, "\n")

	// publish only through pipeline
	if !isDev {
		return container.Publish(ctx, target)
	}

	return
}
