package main

import (
	"context"
	"dagger/ansible/internal/dagger"
	"encoding/json"
	"fmt"
)

type ProductFormat struct {
	Architecture  string `json:"arch"`
	Index         int    `json:"index"`
	OS            string `json:"os"`
	Runner        string `json:"runner"`
	TargetImage   string `json:"target_image"` // without architecture suffix
	UpstreamImage string `json:"upstream_image"`
}

// Product returns the cartesian product of all builds
// this is used to explode the builds
func (m *Ansible) Product(
	ctx context.Context,
	src *dagger.Directory,
	// +default="latest"
	version string,
) (products []ProductFormat, err error) {
	c, err := loadConfig(ctx, src)
	if err != nil {
		return
	}

	// create a list of products
	var i int
	for _, b := range c.Builds {
		for _, a := range b.Architectures {
			runner, err := dag.Lib().ArchToRunner(ctx, a)
			if err != nil {
				return products, err
			}
			products = append(products, ProductFormat{
				Architecture:  a,
				Index:         i,
				OS:            b.OS,
				Runner:        runner,
				TargetImage:   imageTag(c, b, version),
				UpstreamImage: fmt.Sprintf("%s:%s-%s", c.Upstream.Repo, c.Upstream.Tag, b.OS),
			})
			i++
		}
	}

	return
}

func imageTag(c Config, b Build, version string) string {
	return fmt.Sprintf("%s:%s-%s", c.TargetRepo, version, b.OS)
}

// ProductJson returns the cartesian product of all builds as a json string, used for github actions matrix
func (m *Ansible) ProductJson(
	ctx context.Context,
	src *dagger.Directory,
	// +default="latest"
	version string,
) (o string, err error) {
	products, err := m.Product(ctx, src, version)
	bytes, err := json.Marshal(products)
	if err != nil {
		return
	}
	o = string(bytes)
	return
}
