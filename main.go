package main

import (
	"fmt"
	"github.com/sn3d/toolbx/sdk"
	"os"
)

// version is set by goreleaser, via -ldflags="-X 'main.version=...'".
var version = "development"

func main() {

	err := sdk.RunToolbx(
		sdk.WithBrandLabel("toolbx"),
		sdk.WithGitlab(os.Getenv("GITLAB_TOKEN")),
		sdk.WithXdgData(),
		sdk.WithXdgConfig(),
	)

	if err != nil {
		fmt.Errorf("error %v", err)
		os.Exit(1)
	}
}
