package tool

import (
	"github.com/sn3d/toolbx/pkg/command"
	"os"
	"testing"
)

func Test_SaveAndLoadInstallation(t *testing.T) {

	installationsPath, _ := os.MkdirTemp("", "toolbx-installations-*")
	repo := CreateInstallationsRepository(installationsPath)

	// save installed tool
	toolInstance := ToolInstance{
		ID:               "hello",
		InstalledVersion: "1.0.0",
	}

	err := repo.SaveTool(&toolInstance)
	if err != nil {
		t.FailNow()
	}

	// get the installation
	cmd := command.CommandInstance{
		Name:     "hello",
		Metadata: &command.Metadata{},
	}

	savedTool := repo.GetToolForCommand(&cmd)
	if savedTool == nil {
		t.FailNow()
	}

	if savedTool.InstalledVersion != "1.0.0" {
		t.FailNow()
	}

	if savedTool.ID != "hello" {
		t.FailNow()
	}
}
