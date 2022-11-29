package command

import (
	"github.com/sn3d/toolbx/pkg/tempfs"
	"strings"
	"testing"
)

func Test_GetCommand(t *testing.T) {

	for _, table := range []struct {
		description   string
		args          []string
		expectedDir   string
		expectedGroup bool
	}{
		{
			description:   "happy-path",
			args:          []string{"storage", "postgres", "arg1"},
			expectedDir:   "/storage/postgres",
			expectedGroup: false,
		},
		{
			description:   "only-existing",
			args:          []string{"storage", "postgres", "list"},
			expectedDir:   "/storage/postgres",
			expectedGroup: true,
		},
	} {
		t.Run(table.description, func(t *testing.T) {
			dir, err := tempfs.New("./testdata/commands")
			if err != nil {
				t.FailNow()
			}

			cmdRepo := CreateCommandsRepository(dir.GetRoot())
			cmd, err := cmdRepo.GetCommand(table.args)
			if err != nil {
				t.FailNow()
			}

			if !strings.HasSuffix(cmd.Dir, table.expectedDir) {
				t.FailNow()
			}

			if cmd.Metadata == nil {
				t.FailNow()
			}

		})
	}
}

func Test_GetSubcommands(t *testing.T) {

	commands := CreateCommandsRepository("./testdata/commands")

	cmd, err := commands.GetCommand([]string{"cluster"})
	if err != nil {
		t.FailNow()
	}

	subCmds, err := commands.GetSubcommands(cmd)
	if err != nil {
		t.FailNow()
	}

	if len(subCmds) != 2 {
		t.FailNow()
	}
}

// Scenario: load metadata for 'c' subcommand
//    Given: command 'c' with 'command.yaml'
//    When: we load metadata for 'a' command
//    Then: metadata are correctly loaded
//     and: description isn't empty
func Test_loadMetadata(t *testing.T) {

	cmd := CommandInstance{
		Dir: "./testdata/commands/k8s/create",
	}

	err := loadMetadata(&cmd)
	if err != nil {
		t.FailNow()
	}

	if cmd.Metadata.Description == "" {
		t.FailNow()
	}

	if len(cmd.Metadata.Packages) != 3 {
		t.FailNow()
	}
}
