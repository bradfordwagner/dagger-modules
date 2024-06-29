package main

import (
	"context"
	"gopkg.in/yaml.v3"
)

func (m *Ansible) Init(
	ctx context.Context,
	src *Directory,
) (s string, err error) {
	// default config
	c := Config{
		TargetRepo: "ghcr.io/bradfordwagner/gh-template-ansible-playbook",
		Upstream: Upstream{
			Repo: "ghcr.io/bradfordwagner/ansible",
			Tag:  "5.9.1",
		},
		Builds: []Build{
			Build{OS: "alpine_3.18", Architectures: []string{"linux/amd64", "linux/arm64"}},
			Build{OS: "alpine_3.19", Architectures: []string{"linux/amd64", "linux/arm64"}},
			Build{OS: "alpine_3.20", Architectures: []string{"linux/amd64", "linux/arm64"}},
			Build{OS: "archlinux_latest", Architectures: []string{"linux/amd64"}},
			Build{OS: "debian_bookworm", Architectures: []string{"linux/amd64", "linux/arm64"}},
			Build{OS: "debian_bookworm-slim", Architectures: []string{"linux/amd64", "linux/arm64"}},
			Build{OS: "rockylinux_8", Architectures: []string{"linux/amd64", "linux/arm64"}},
			Build{OS: "rockylinux_9", Architectures: []string{"linux/amd64", "linux/arm64"}},
			Build{OS: "ubuntu_jammy", Architectures: []string{"linux/amd64", "linux/arm64"}},
			Build{OS: "ubuntu_mantic", Architectures: []string{"linux/amd64", "linux/arm64"}},
			Build{OS: "ubuntu_noble", Architectures: []string{"linux/amd64", "linux/arm64"}},
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
