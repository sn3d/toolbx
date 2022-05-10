package toolbx

import (
	"os"
	"path/filepath"
	"testing"
)

func Test_Sync(t *testing.T) {
	repo := os.Getenv("TOOLBXREPO")
	if repo == "" {
		t.Skip("set TOOLBXREPO to run this testutil")
	}

	tempDir, err := os.MkdirTemp("", "toolbx-sync")
	if err != nil {
		t.FailNow()
	}

	t.Run("happy-path", func(t *testing.T) {
		err = sync(repo, "main", "", tempDir, filepath.Join(tempDir, "commands"))
		if err != nil {
			t.FailNow()
		}
	})
}
