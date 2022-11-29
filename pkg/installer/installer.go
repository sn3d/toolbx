package installer

import (
	"fmt"
	"github.com/sn3d/toolbx/pkg/installer/archive"
	"net/url"
)

type Installer interface {
	Install(uri url.URL, destDir string) error
}

type InstallationOptions struct {
	BearerToken string
}

func Install(uri url.URL, destDir string, opts InstallationOptions) error {
	if uri.Scheme == "https+zip" {
		return archive.Install(uri, destDir, opts.BearerToken)
	} else {
		return fmt.Errorf("unsupported package scheme %s", uri.Scheme)
	}
}
