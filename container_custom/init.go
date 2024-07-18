package main

import (
	"context"
	"dagger/container-custom/internal/dagger"

	"gopkg.in/yaml.v3"
)

// Init creates an example yaml config for cicd to use
func (m *ContainerCustom) Init(
	ctx context.Context,
	src *dagger.Directory,
) (s string, err error) {
	// default config
	c := Config{
		TargetRepo: "ghcr.io/bradfordwagner/ansible",
		Upstream: Upstream{
			Repo: "ghcr.io/bradfordwagner/base",
			Tag:  "3.6.0",
		},
		Builds: []Build{
			Build{OS: "alpine_3.18", Architectures: []string{"linux/amd64", "linux/arm64"}, Args: map[string]string{"pkg_installer": "alpine"}},
			Build{OS: "alpine_3.19", Architectures: []string{"linux/amd64", "linux/arm64"}, Args: map[string]string{"pkg_installer": "alpine"}},
			Build{OS: "alpine_3.20", Architectures: []string{"linux/amd64", "linux/arm64"}, Args: map[string]string{"pkg_installer": "alpine"}},
			Build{OS: "archlinux_latest", Architectures: []string{"linux/amd64"}, Args: map[string]string{"pkg_installer": "arch"}},
			Build{OS: "debian_bookworm", Architectures: []string{"linux/amd64", "linux/arm64"}, Args: map[string]string{"pkg_installer": "debian"}},
			Build{OS: "debian_bookworm-slim", Architectures: []string{"linux/amd64", "linux/arm64"}, Args: map[string]string{"pkg_installer": "debian"}},
			Build{OS: "rockylinux_8", Architectures: []string{"linux/amd64", "linux/arm64"}, Args: map[string]string{"pkg_installer": "rhel"}},
			Build{OS: "rockylinux_9", Architectures: []string{"linux/amd64", "linux/arm64"}, Args: map[string]string{"pkg_installer": "rhel"}},
			Build{OS: "ubuntu_jammy", Architectures: []string{"linux/amd64", "linux/arm64"}, Args: map[string]string{"pkg_installer": "debian"}},
			Build{OS: "ubuntu_mantic", Architectures: []string{"linux/amd64", "linux/arm64"}, Args: map[string]string{"pkg_installer": "debian"}},
			Build{OS: "ubuntu_noble", Architectures: []string{"linux/amd64", "linux/arm64"}, Args: map[string]string{"pkg_installer": "debian"}},
		},
	}

	// convert to yaml
	b, err := yaml.Marshal(c)
	if err != nil {
		return
	}

	// return yaml
	return string(b), nil
}
