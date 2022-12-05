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
		sdk.WithXdg("toolbx"), // configuration is loaded from $HOME/.config/toolbx/toolbx.yaml and data are stored in $HOME/.local/share/toolbx
		sdk.WithBearerToken(os.Getenv("GITLAB_TOKEN")),
	)

	if err != nil {
		fmt.Errorf("error %v", err)
		os.Exit(1)
	}
}
