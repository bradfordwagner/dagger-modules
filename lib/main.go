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
	// GitHub actor, --token=env:GITHUB_API_TOKEN,--token=cmd:"gh auth token"
	actor *Secret,
	// GitHub API token, --token=env:GITHUB_API_TOKEN,--token=cmd:"gh auth token"
	token *Secret,
	image string,
	arches []string,
) (s string, err error) {
	// --platforms linux/amd64,linux/s390x,linux/arm64 \
	// --template foo/bar-ARCH:v1 \
	// --target foo/bar:v1
	// return dag.Container().From("mplatform/manifest-tool:alpine-v2.1.6").
	return dag.Container().From("mplatform/manifest-tool:alpine-v2.0.3").
		WithSecretVariable("GITHUB_ACTOR", actor).
		WithSecretVariable("GITHUB_TOKEN", token).
		WithFocus().
		WithExec([]string{
			"--username", "${GITHUB_ACTOR}",
			"--password", "${GITHUB_TOKEN}",
			"push", "from-args",
			"--platforms", strings.Join(arches, ","),
			"--template", image + "-OS_ARCH",
			"--target", image,
		}).Stderr(ctx)
}
