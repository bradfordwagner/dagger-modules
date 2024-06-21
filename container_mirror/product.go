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
	"encoding/json"
)

type Product struct {
	Repo         string `json:"repo"`
	Tag          string `json:"tag"`
	Architecture string `json:"arch"`
}

// Cartesian returns the cartesian product of all builds
// this is used to explode the builds
func (m *ContainerMirror) Product(
	ctx context.Context,
	src *Directory,
) (o string, err error) {
	c, err := loadConfig(ctx, src)
	if err != nil {
		return
	}

	// create a list of products
	var products []Product
	for _, b := range c.Builds {
		for _, a := range b.Architectures {
			products = append(products, Product{
				Repo:         b.Repo,
				Tag:          b.Tag,
				Architecture: a,
			})
		}
	}

	// convert products to json
	bytes, err := json.Marshal(products)
	o = string(bytes)
	return
}
