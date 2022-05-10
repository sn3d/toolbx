package api

import (
	"path/filepath"
)

type Installation struct {
	ID               string `yaml:"id"`
	InstalledVersion string `yaml:"installedVersion"`
	InstalledCmd     string `yaml:"installedCmd"`
}

func (i *Installation) Dir(rootDir string) string {
	installationDir := filepath.Join(rootDir, i.ID, i.InstalledVersion)
	return installationDir
}

func (i *Installation) Binary(rootDir string) string {
	return filepath.Join(i.Dir(rootDir), i.InstalledCmd)
}
