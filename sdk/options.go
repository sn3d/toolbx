package sdk

import (
	"github.com/sn3d/toolbx/pkg/config"
	"github.com/sn3d/toolbx/pkg/dir"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
)

type ToolbxOption func(cfg *config.Configuration)

// WithXdg ensure the toolbx will use XDG directory spec.
// for config files and all data files.
//
// The data will be stored in $HOME/.local/share/{name} and configuration
// will be loaded from $HOME/.config/{name}/{name}.yaml
//
func WithXdg(name string) ToolbxOption {
	return func(cfg *config.Configuration) {
		configHome := dir.XdgConfigHome()
		toolbxConfigFile := path.Join(configHome, cfg.BrandLabel, cfg.BrandLabel+".yaml")
		WithConfigFile(toolbxConfigFile)(cfg)

		dataHome := dir.XdgDataHome()
		toolbxDataHome := path.Join(dataHome, cfg.BrandLabel)
		WithDataDir(toolbxDataHome)(cfg)
	}
}

// WithConfigFile ensure the configuration will be loaded
// from path. Path need to contain also YAML file.
func WithConfigFile(path ...string) ToolbxOption {
	return func(cfg *config.Configuration) {
		cfgFile := filepath.Join(path...)
		if _, err := os.Stat(cfgFile); os.IsNotExist(err) {
			return
		}

		yamlFile, err := ioutil.ReadFile(cfgFile)
		if err != nil {
			return
		}

		var loadedCfg config.Configuration
		err = yaml.Unmarshal(yamlFile, &loadedCfg)
		if err != nil {
			return
		}

		if loadedCfg.CommandsRepository != "" {
			cfg.CommandsRepository = loadedCfg.CommandsRepository
		}

		if loadedCfg.CommandsBranch != "" {
			cfg.CommandsBranch = loadedCfg.CommandsBranch
		}
	}
}

// If you want to use GitLab for distribution, you need to
// provide GitLab personal access token
func WithBearerToken(token string) ToolbxOption {
	return func(cfg *config.Configuration) {
		cfg.Token = token
	}
}

func WithCommandsRepository(repo string, branch string) ToolbxOption {
	return func(cfg *config.Configuration) {
		cfg.CommandsRepository = repo
		if branch != "" {
			cfg.CommandsBranch = branch
		}
	}
}

// WithDataDir set data directory, where are placed
// installations etc...
func WithDataDir(path ...string) ToolbxOption {
	return func(cfg *config.Configuration) {
		dataDir := filepath.Join(path...)
		if dataDir != "" {
			dir.Ensure(dataDir)
			cfg.DataDir = dataDir
		}
	}
}

func WithBrandLabel(brand string) ToolbxOption {
	return func(cfg *config.Configuration) {
		cfg.BrandLabel = brand
	}
}
