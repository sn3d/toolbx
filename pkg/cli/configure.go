package cli

import (
	"flag"
	"fmt"
	"github.com/sn3d/toolbx/pkg/config"
	"github.com/sn3d/toolbx/pkg/dir"

	"gopkg.in/yaml.v3"
	"io/ioutil"
	"path"
)

// ConfigureCmd is one of the dot commands, and it's executed when you
// type `toolbx .configure`
// This command run configuration wizard that helps you setup
// new installation of toolbx
func ConfigureCmd(args []string) {
	config := &config.Configuration{}

	fs := flag.NewFlagSet("configure", flag.ContinueOnError)
	fs.StringVar(&config.CommandsRepository, "repository", "", "URL to repository with commands definitions (e.g. https://github.com/sn3d/toolbx-demo.git)")
	fs.StringVar(&config.CommandsBranch, "branch", "main", "branch that will be used (default is main)")

	fs.Parse(args)

	configDir := path.Join(dir.XdgConfigHome(), "toolbx")
	dir.Ensure(configDir)

	configFile := path.Join(configDir, "toolbx.yaml")
	err := saveConfig(config, configFile)
	if err != nil {
		fmt.Printf("Error: %w", err)
	}
}

func saveConfig(cfg *config.Configuration, file string) error {
	data, err := yaml.Marshal(cfg)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(file, data, 0664)
	return err
}
