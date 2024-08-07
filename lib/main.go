// A generated module for Lib functions
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
	"strings"
	"time"
)

type Lib struct{}

// Returns lines that match a pattern in the files of the provided Directory
func (m *Lib) OpenConfigYaml(ctx context.Context, src Directory) (s string, err error) {
	configFile := src.File("config.yaml")
	return configFile.Contents(ctx)
}

// Returns the runner to run on for the given architecture
func (m *Lib) ArchToRunner(arch string) (s string) {
	archs := map[string]string{
		"linux/arm64": "arm64",
	}

	// default to ubuntu-latest
	s, ok := archs[arch]
	if !ok {
		return "ubuntu-latest"
	}
	return
}

// ArchImageName returns the image name for the given image and architecture
func (m *Lib) ArchImageName(image, arch string) (s string) {
	arch = strings.ReplaceAll(arch, "/", "_")
	return image + "-" + arch
}

func (m *Lib) ManifestTool(
	ctx context.Context,
	// GitHub actor, --token=env:github_actor,--token=cmd:"gh auth token"
	actor *Secret,
	// GitHub API token, --token=env:github_token,--token=cmd:"gh auth token"
	token *Secret,
	image string,
	arches []string,
) (s string, err error) {
	username, _ := actor.Plaintext(ctx)
	password, _ := token.Plaintext(ctx)
	c := dag.Container().From("mplatform/manifest-tool:alpine-v2.1.6").
		WithFocus().
		WithExec([]string{
			"--username", username,
			"--password", password,
			"push", "from-args",
			"--platforms", strings.Join(arches, ","),
			"--template", image + "-OS_ARCH",
			"--target", image,
		})
	return m.ContainerOutput(ctx, c)
}

// ContainerOutput returns the output of a container as a string if stderr exists return that, else stdout
func (m *Lib) ContainerOutput(ctx context.Context, c *Container) (s string, err error) {
	s, err = c.Stderr(ctx)
	if err != nil || s != "" {
		return
	}
	return c.Stdout(ctx)
}

// FileContents - returns the contents of a file in a directory
func (m *Lib) FileContents(ctx context.Context, dir *Directory, path string) (contents string, err error) {
	if file := dir.File(path); file != nil {
		contents, err = file.Contents(ctx)
	}
	return
}

// InvalidateCache invalidates the cache if the shouldInvalidate is true
func (m *Lib) InvalidateCache(shouldInvalidate bool, container *Container) *Container {
	if shouldInvalidate {
		container = container.WithEnvVariable("CACHEBUSTER", time.Now().String())
	}
	return container
}
