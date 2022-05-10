package api

import (
	"testing"
)

func Test_InstallationID(t *testing.T) {

	// Scenario: resolve installation ID hierarchically by joining all parents
	// names and name of the current name.
	//
	// It's very same logic the Git is using e.g. for subcommand 'one two three'
	// the binary will be 'one-two-three'.
	t.Run("binaryAsChain", func(t *testing.T) {
		level1Cmd := &Command{
			Name:     "a",
			Metadata: &Metadata{},
		}

		level2Cmd := &Command{
			Name:     "b",
			Parent:   level1Cmd,
			Metadata: &Metadata{},
		}

		level3Cmd := &Command{
			Name:     "c",
			Parent:   level2Cmd,
			Metadata: &Metadata{},
		}

		binary := level3Cmd.GetInstallationID()
		if binary != "a-b-c" {
			t.FailNow()
		}
	})

	// Scenario: when root parent have empty name
	t.Run("rootWithoutName", func(t *testing.T) {
		level1Cmd := &Command{
			Name: "",
		}

		level2Cmd := &Command{
			Name:   "hello",
			Parent: level1Cmd,
		}

		level3Cmd := &Command{
			Name:   "subcommand",
			Parent: level2Cmd,
		}

		binary := level3Cmd.GetInstallationID()
		if binary != "hello-subcommand" {
			t.FailNow()
		}
	})
}
