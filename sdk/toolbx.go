package sdk

import (
	"fmt"
	"github.com/sn3d/toolbx/pkg/config"
	"github.com/sn3d/toolbx/pkg/executor"
	"os"
)

func RunToolbx(options ...ToolbxOption) error {
	cfg := config.Configuration{}

	for _, option := range options {
		option(&cfg)
	}

	exec, err := executor.Initialize(cfg)
	if err != nil {
		fmt.Errorf("error %v", err)
		os.Exit(1)
	}

	err = exec.Execute(os.Args)
	if err != nil {
		fmt.Errorf("error %v", err)
		os.Exit(1)
	}

	return nil
}
