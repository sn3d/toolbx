package toolbx

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"path/filepath"
	"toolbx/api"
)

type InstallationsRepository struct {
	InstallationsDir string
}

func CreateInstallationsRepository(installationDir string) *InstallationsRepository {
	return &InstallationsRepository{
		InstallationsDir: installationDir,
	}
}

func (fsi *InstallationsRepository) GetInstallationForCommand(cmd *api.Command) *api.Installation {
	id := cmd.GetInstallationID()
	path := filepath.Join(fsi.InstallationsDir, id+".yaml")

	installation := &api.Installation{}
	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		return nil

	}

	err = yaml.Unmarshal(yamlFile, installation)
	if err != nil {
		return nil
	}

	installation.ID = id
	return installation
}

func (fsi *InstallationsRepository) SaveInstallation(i *api.Installation) error {
	path := filepath.Join(fsi.InstallationsDir, i.ID+".yaml")
	data, err := yaml.Marshal(i)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(path, data, 0640)
	if err != nil {
		return err
	}
	return nil
}
