package tool

import (
	"path/filepath"
)

// Tool represent installed tool in your system. Tool is installed
// from Package and executed from Command
type ToolInstance struct {
	ID               string `yaml:"id"`
	InstalledVersion string `yaml:"installedVersion"`
	InstalledCmd     string `yaml:"installedCmd"`
}

func (i *ToolInstance) Dir(rootDir string) string {
	installationDir := filepath.Join(rootDir, i.ID, i.InstalledVersion)
	return installationDir
}

func (i *ToolInstance) Binary(rootDir string) string {
	return filepath.Join(i.Dir(rootDir), i.InstalledCmd)
}
