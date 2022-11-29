package command

import (
	"runtime"
	"strings"
)

type CommandInstance struct {
	Parent   *CommandInstance
	Name     string
	Dir      string
	Args     []string
	Metadata *Metadata
}

type CommandMetadata struct {
	Version     string    `yaml:"version"`
	Description string    `yaml:"description",omitempty`
	Usage       string    `yaml:"usage",omitempty`
	Packages    []Package `yaml:"packages"`
}

// GetToolID returns you tool's ID in form of a flat-name
// with all parents e.g. for command 'hello world' it
// returns 'hello-world'
func (cmd *CommandInstance) GetToolID() string {
	var prefix = ""
	if cmd.Parent != nil {
		if cmd.Parent.Name != "" {
			prefix = cmd.Parent.GetToolID() + "-"
		}
	}
	return prefix + cmd.Name
}

// GetPackage returns you package for your OS and ARCH. If there
// is no package for your platform available, then it returns nil
func (cmd *CommandInstance) GetPackage() *Package {
	platform := strings.ToLower(runtime.GOOS + "-" + runtime.GOARCH)
	for _, pkg := range cmd.Metadata.Packages {
		if strings.ToLower(pkg.Platform) == platform {
			return &pkg
		}
	}
	return nil
}
