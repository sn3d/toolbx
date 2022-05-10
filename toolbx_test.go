package toolbx

import (
	"os"
	"testing"
)

func Test_Execute(t *testing.T) {
	repo := os.Getenv("TOOLBXREPO")
	if repo == "" {
		t.Skip("set TOOLBXREPO to demo repository if you want to run this testutil")
	}

	toolbxpath, err := os.MkdirTemp("", "toolbx-exec")
	if err != nil {
		t.FailNow()
	}

	toolbx, err := Create(
		WithToolbxPath(toolbxpath),
		WithSyncRepo(os.Getenv("TOOLBXREPO"), "main"),
	)

	if err != nil {
		t.FailNow()
	}

	t.Run("simple-test", func(t *testing.T) {
		args := []string{"toolbx", "k8s", "create", "arg1", "arg2"}
		err = toolbx.Execute(args)
		if err != nil {
			t.FailNow()
		}
	})

	t.Run("empty-subcommands", func(t *testing.T) {
		args := []string{"toolbx"}
		err = toolbx.Execute(args)
		if err != nil {
			t.FailNow()
		}
	})

	t.Run("group", func(t *testing.T) {
		args := []string{"toolbx", "k8s"}
		err = toolbx.Execute(args)
		if err != nil {
			t.FailNow()
		}
	})

}
