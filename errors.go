package toolbx

import "errors"

var (
	// NoChildError signalising there is no child command of parent
	// command
	NoChildError = errors.New("no child subcommands")

	// NoMetadataError indicates the command.yaml is missing in the subcommand
	// folder
	NoMetadataError = errors.New("no metadata file")
)
