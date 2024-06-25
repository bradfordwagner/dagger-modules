package main

import (
	"context"

	"gopkg.in/yaml.v3"
)

type ContainerCustom struct{}

type Config struct {
	TargetRepo string   `yaml:"target_repo"`
	Builds     []Build  `yaml:"builds"`
	Upstream   Upstream `yaml:"upstream"`
}

type Upstream struct {
	Repo string `yaml:"repo"`
	Tag  string `yaml:"tag"`
}

type Build struct {
	OS            string            `yaml:"os"`
	Architectures []string          `yaml:"archs"`
	Args          map[string]string `yaml:"args"`
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
