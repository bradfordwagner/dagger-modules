package main

import (
	"context"
	"dagger/ansible/internal/dagger"
	"encoding/json"
	"strings"
)

// Build - builds the container image
func (m *Ansible) Build(
	ctx context.Context,
	// +default=0
	index int,
	src *Directory,
	// +default="latest"
	version string,
	// +default=false
	invalidateCache bool,
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

	// build the container
	container := dag.Container(ContainerOpts{
		Platform: dagger.Platform(product.Architecture),
	}).From(product.UpstreamImage).WithDirectory("/src", src)
	container = dag.Lib().InvalidateCache(invalidateCache, container)

	// find requirements
	requirements := []string{"requirements.yml", "meta/requirements.yml"}
	dir := container.Directory("/src")
	for _, requirement := range requirements {
		var contents string
		if contents, err = dag.Lib().FileContents(ctx, dir, requirement); contents != "" {
			container, err = container.WithExec([]string{"ansible-galaxy", "install", "-r", requirement}).Sync(ctx)
			if err != nil {
				return
			}
		}
	}

	// run playbook
	playbooks := []string{"test.yml", "playbook.yml"}
	for _, playbook := range playbooks {
		var contents string
		if contents, err = dag.Lib().FileContents(ctx, dir, playbook); contents != "" {
			container, err = container.WithExec([]string{"ansible-playbook", playbook}).Sync(ctx)
			if err != nil {
				return
			}
		}
	}
	// zero error after execution to allow for missing playbook entries
	err = nil
	o = strings.Join([]string{target, productJson}, "\n")

	// publish only through pipeline
	if !isDev {
		_, err = container.Publish(ctx, target)
	}

	return
}
