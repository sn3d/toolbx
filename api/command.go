package api

import (
	"runtime"
	"strings"
)

type Command struct {
	Parent   *Command
	Name     string
	Dir      string
	Args     []string
	Metadata *Metadata
}

// GetInstallationID returns you identifier in form of a flat
// name with all parents e.g. for command 'hello world' it
// returns 'hello-world'
func (cmd *Command) GetInstallationID() string {
	var prefix = ""
	if cmd.Parent != nil {
		if cmd.Parent.Name != "" {
			prefix = cmd.Parent.GetInstallationID() + "-"
		}
	}
	return prefix + cmd.Name
}

// GetPackage returns you package for your OS and ARCH. If there
// is no package for your platform available, then it returns nil
func (cmd *Command) GetPackage() *Package {
	platform := strings.ToLower(runtime.GOOS + "-" + runtime.GOARCH)
	for _, pkg := range cmd.Metadata.Packages {
		if strings.ToLower(pkg.Platform) == platform {
			return &pkg
		}
	}
	return nil
}
