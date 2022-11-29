package dir

import (
	"os"
	"path"
)

func UserHome() string {
	homeDir, _ := os.UserHomeDir()
	return homeDir
}

func XdgDataHome() string {
	xdgConfigHome := os.Getenv("XDG_DATA_HOME")
	if xdgConfigHome == "" {
		xdgConfigHome = path.Join(UserHome(), ".local", "share")
	}
	return xdgConfigHome

}

func XdgConfigHome() string {
	xdgConfigHome := os.Getenv("XDG_CONFIG_HOME")
	if xdgConfigHome == "" {
		xdgConfigHome = path.Join(UserHome(), ".config")
	}
	return xdgConfigHome
}

func Ensure(dirElms ...string) string {
	dir := path.Join(dirElms...)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		os.MkdirAll(dir, os.ModePerm)
	}
	return dir
}
