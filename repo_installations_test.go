package toolbx

import (
	"os"
	"testing"
	"toolbx/api"
)

func Test_SaveAndLoadInstallation(t *testing.T) {

	installationsPath, _ := os.MkdirTemp("", "toolbx-installations-*")
	repo := CreateInstallationsRepository(installationsPath)

	// save installation
	inst := api.Installation{
		ID:               "hello",
		InstalledVersion: "1.0.0",
	}

	err := repo.SaveInstallation(&inst)
	if err != nil {
		t.FailNow()
	}

	// get the installation
	cmd := api.Command{
		Name:     "hello",
		Metadata: &api.Metadata{},
	}

	savedInst := repo.GetInstallationForCommand(&cmd)
	if savedInst == nil {
		t.FailNow()
	}

	if savedInst.InstalledVersion != "1.0.0" {
		t.FailNow()
	}

	if savedInst.ID != "hello" {
		t.FailNow()
	}
}
