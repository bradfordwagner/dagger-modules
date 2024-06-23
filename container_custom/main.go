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
	RepoOverride  string   `yaml:"repo_override"` // renames repo to override in the target image
	Tag           string   `yaml:"tag"`
	Architectures []string `yaml:"archs"`
}

// loadConfig loads the config.yaml from the source directory
func loadConfig(ctx context.Context, src *Directory) (c Config, err error) {
	yml, _ := dag.Lib().OpenConfigYaml(ctx, src)
	err = yaml.Unmarshal([]byte(yml), &c)
	if err != nil {
		return
	}
	return
}
