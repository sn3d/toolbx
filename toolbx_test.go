package toolbx

import (
	"testing"
	"toolbx/testutil"
)

func Test_Execute(t *testing.T) {
	toolbxpath, err := testutil.CreateTestData("./testdata")
	if err != nil {
		t.FailNow()
	}

	toolbx, err := Create(
		WithToolbxPath(toolbxpath),
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
