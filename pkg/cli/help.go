package cli

import (
	"fmt"
	"github.com/fatih/color"
)

// HelpCmd is behind '.help' dot command and print the list
// of the all dot commands
func HelpCmd(args []string) {
	bold := color.New(color.FgHiWhite, color.Bold).SprintfFunc()

	fmt.Printf("List of dot-commands:\n")
	fmt.Printf("\n")
	fmt.Printf("   %s: %s\n", bold("configure"), "configure/reconfigure toolbx on your machine")
	fmt.Printf("   %s: %s\n", bold("help"), "print this help")
	fmt.Printf("\n")
}
