package install

import "net/url"

type Installer interface {
	Install(uri url.URL, destDir string) error
}
