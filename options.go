package toolbx

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
	"path/filepath"
)

type ToolbxOption func(instance *Toolbx)

func defaultValues(instance *Toolbx) {
	instance.syncRepoBranch = "main"

	WithNameFromArgs()
	WithToolbxPath(os.Getenv("TOOLBXPATH"))(instance)
	WithGitlab(os.Getenv("GITLAB_TOKEN"))(instance)
}

func WithToolbxPath(toolbxpath string) ToolbxOption {
	toolbxpath = toolbxPath(toolbxpath)
	homeDir, _ := os.UserHomeDir()
	return func(instance *Toolbx) {
		instance.syncFile = filepath.Join(toolbxpath, "sync")
		WithInstallationsDir(filepath.Join(toolbxpath, "installed"))(instance)
		WithCommandsDir(filepath.Join(toolbxpath, "commands"))(instance)

		// The configuration is loaded in following order:
		//    - ~/.toolbx.yaml
		//    - $TOOLBX/toolbx.yaml
		WithConfigFile(filepath.Join(homeDir, ".toolbx.yaml"))(instance)
		WithConfigFile(filepath.Join(toolbxpath, "toolbx.yaml"))(instance)
	}
}

func WithConfigFile(path string) ToolbxOption {
	return func(instance *Toolbx) {
		if _, err := os.Stat(path); os.IsNotExist(err) {
			return
		}

		yamlFile, err := ioutil.ReadFile(path)
		if err != nil {
			return
		}

		var config Config
		err = yaml.Unmarshal(yamlFile, &config)
		if err != nil {
			return
		}

		if config.InstallationsPath != "" {
			instance.installationsDir = config.InstallationsPath
		}

		if config.CommandsPath != "" {
			instance.commandsDir = config.CommandsPath
		}

		if config.Repository != "" {
			instance.syncRepo = config.Repository
		}

		if config.Branch != "" {
			instance.syncRepoBranch = config.Branch
		}

		if config.SyncFile != "" {
			instance.syncFile = config.SyncFile
		}
	}
}

func WithGitlab(personalAccessToken string) ToolbxOption {
	return func(instance *Toolbx) {
		instance.gitlabToken = personalAccessToken
	}
}

func WithSyncRepo(repo string, branch string) ToolbxOption {
	return func(instance *Toolbx) {
		instance.syncRepo = repo
		if branch != "" {
			instance.syncRepoBranch = branch
		}
	}
}

func WithInstallationsDir(dir string) ToolbxOption {
	return func(instance *Toolbx) {
		if dir != "" {
			instance.installationsDir = dir
		}
	}
}

func WithCommandsDir(dir string) ToolbxOption {
	return func(instance *Toolbx) {
		if dir != "" {
			instance.commandsDir = dir
		}
	}
}

func WithName(name string) ToolbxOption {
	return func(instance *Toolbx) {
		instance.name = name
	}
}

func WithNameFromArgs() ToolbxOption {
	base := filepath.Base(os.Args[0])
	return func(instance *Toolbx) {
		instance.name = base
	}
}

func toolbxPath(path string) string {
	if path == "" {
		home, _ := os.UserHomeDir()
		return filepath.Join(home, ".toolbx")
	}
	return path
}
