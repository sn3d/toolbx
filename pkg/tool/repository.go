package tool

import (
	"github.com/sn3d/toolbx/pkg/command"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"path/filepath"
)

type InstalledToolsRepository struct {
	// directory where are tools installed on your system
	InstallationsDir string
}

func CreateInstallationsRepository(InstallationsDir string) *InstalledToolsRepository {
	return &InstalledToolsRepository{
		InstallationsDir: InstallationsDir,
	}
}

func (fsi *InstalledToolsRepository) GetToolForCommand(cmd *command.CommandInstance) *ToolInstance {
	id := cmd.GetToolID()
	path := filepath.Join(fsi.InstallationsDir, id+".yaml")

	installedTool := &ToolInstance{}
	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		return nil

	}

	err = yaml.Unmarshal(yamlFile, installedTool)
	if err != nil {
		return nil
	}

	installedTool.ID = id
	return installedTool
}

func (fsi *InstalledToolsRepository) SaveTool(t *ToolInstance) error {
	path := filepath.Join(fsi.InstallationsDir, t.ID+".yaml")
	data, err := yaml.Marshal(t)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(path, data, 0640)
	if err != nil {
		return err
	}
	return nil
}
