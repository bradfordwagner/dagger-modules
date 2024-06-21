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

	"gopkg.in/yaml.v2"
)

// Init creates an example yaml config for cicd to use
func (m *ContainerMirror) Init(
	ctx context.Context,
	src *Directory,
) (err error) {

	c := Config{
		TargetRepo: "ghcr.io/bradfordwagner/template-mirror",
		Builds: []Build{
			{Repo: "alpine", Tag: "3.19", Architectures: []string{"linux/amd64", "linux/arm64"}},
			{Repo: "alpine", Tag: "3.18", Architectures: []string{"linux/amd64", "linux/arm64"}},
			{Repo: "debian", Tag: "bookworm", Architectures: []string{"linux/amd64", "linux/arm64"}},
			{Repo: "debian", Tag: "bullseye", Architectures: []string{"linux/amd64", "linux/arm64"}},
			{Repo: "ubuntu", Tag: "noble", Architectures: []string{"linux/amd64", "linux/arm64"}},
			{Repo: "ubuntu", Tag: "mantic", Architectures: []string{"linux/amd64", "linux/arm64"}},
			{Repo: "archlinux", Tag: "latest", Architectures: []string{"linux/amd64"}},
		},
	}

	// src.WithNewFile(path string, contents string, opts ...dagger.DirectoryWithNewFileOpts)
	b, err := yaml.Marshal(c)
	if err != nil {
		return
	}

	src.WithNewFile("config.yaml", string(b), dagger.DirectoryWithNewFileOpts{Permissions: 755})
	return
}
