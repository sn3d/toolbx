package sdk

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/sn3d/toolbx/pkg/config"
	"github.com/sn3d/toolbx/pkg/executor"
	"os"
)

func RunToolbx(options ...ToolbxOption) error {
	cfg := config.Configuration{}

	for _, option := range options {
		option(&cfg)
	}

	exec := executor.Create(cfg)
	err := exec.Execute(os.Args)
	if err != nil {
		red := color.New(color.FgHiRed).SprintfFunc()
		fmt.Printf("%s: %v\n", red("error"), err)
		os.Exit(1)
	}

	return nil
}
