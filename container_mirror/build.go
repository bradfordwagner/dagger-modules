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
	"fmt"
)

func (m *ContainerMirror) Build(
	ctx context.Context,
	// +default=0
	index int,
	src *Directory,
) (o string, err error) {
	products, err := m.Product(ctx, src)
	if err != nil {
		return
	}

	product := products[index]
	b, err := json.Marshal(product)
	if err != nil {
		return
	}
	o = string(b)

	dag.Container().
		From(fmt.Sprintf("%s:%s", product.Repo, product.Tag))
	return
}
