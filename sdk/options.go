package sdk

import (
	"github.com/sn3d/toolbx/pkg/config"
	"github.com/sn3d/toolbx/pkg/dir"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
	"path"
)

type ToolbxOption func(cfg *config.Configuration)

// load configuration file from $HOME/.config/{brand label}/{brand label}.yaml
func WithXdgConfig() ToolbxOption {
	return func(cfg *config.Configuration) {
		configHome := dir.XdgConfigHome()
		toolbxConfigFile := path.Join(configHome, cfg.BrandLabel, cfg.BrandLabel+".yaml")
		WithConfigFile(toolbxConfigFile)(cfg)
	}
}

// data like installations will be stored in $HOME/.local/share/{brand label}
func WithXdgData() ToolbxOption {
	return func(cfg *config.Configuration) {
		dataHome := dir.XdgDataHome()
		toolbxDataHome := path.Join(dataHome, cfg.BrandLabel)
		WithDataDir(toolbxDataHome)(cfg)
	}
}

// WithConfigFile ensure the configuration will be loaded
// from path. Path need to contain also YAML file.
func WithConfigFile(path string) ToolbxOption {
	return func(cfg *config.Configuration) {
		if _, err := os.Stat(path); os.IsNotExist(err) {
			return
		}

		yamlFile, err := ioutil.ReadFile(path)
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
// provide GitLab token
func WithGitlab(personalAccessToken string) ToolbxOption {
	return func(cfg *config.Configuration) {
		cfg.GitlabToken = personalAccessToken
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
func WithDataDir(dataDir string) ToolbxOption {
	return func(cfg *config.Configuration) {
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
