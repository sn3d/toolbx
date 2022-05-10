package toolbx

import "errors"

var (
	// NoChildError signalising there is no child command of parent
	// command
	NoChildError = errors.New("no child subcommands")

	// NoMetadataError indicates the command.yaml is missing in the subcommand
	// folder
	NoMetadataError = errors.New("no metadata file")

	// MissingRepoError indicates you have probably not configured toolbx. You should
	// create configuration file and define Git repository where are commands
	// defined:
	//
	//     echo "repository: https://github.com/sn3d/toolbx-demo.git" > ~/.toolbx.yaml
	//
	MissingRepoError = errors.New("no repository with commands defined")
)
