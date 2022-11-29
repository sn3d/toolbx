package command

import "errors"

// NoChildError signalising there is no child command of parent
// command
var NoChildError = errors.New("no child subcommands")
