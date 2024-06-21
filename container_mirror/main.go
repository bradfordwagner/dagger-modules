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

	"gopkg.in/yaml.v3"
)

type ContainerMirror struct{}

type Config struct {
	TargetRepo string  `yaml:"target_repo"`
	Builds     []Build `yaml:"builds"`
}

type Build struct {
	Repo          string   `yaml:"repo"`
	Tag           string   `yaml:"tag"`
	Architectures []string `yaml:"archs"`
}

// Returns a container that echoes whatever string argument is provided
// func (m *ContainerMirror) Mirror() *Container {
func (m *ContainerMirror) Mirror(
	ctx context.Context,
	src *Directory,
) (o string, err error) {
	c, err := loadConfig(ctx, src)
	if err != nil {
		return
	}

	return c.TargetRepo, nil
}

func loadConfig(ctx context.Context, src *Directory) (c Config, err error) {
	yml, _ := dag.Lib().OpenConfigYaml(ctx, src)
	err = yaml.Unmarshal([]byte(yml), &c)
	if err != nil {
		return
	}
	return
}
