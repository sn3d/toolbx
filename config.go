package toolbx

type Config struct {
	// path to all installed binaries e.g. $TOOLBX/installations
	InstallationsPath string `yaml:"installationsPath",omitempty`

	// path to all command definitions, it's synced with Git repository (e.g. $TOOLBX/commands)
	CommandsPath string `yaml:"commandsPath",omitempty`

	// URL to Git repository with commands you want to sync with (e.g. https://gitlab.myorg.com/toolbx-commands.git
	Repository string `yaml:"repository",omitempty`

	// What to say, branch that will be used for sync
	Branch string `yaml:"branch",omitempty`

	// path to sync file. It's empty file that serve for ensuring
	// sync is executed only once per day ( e.g. $TOOLBX/sync )
	SyncFile string `yaml:"syncFile",omitempty`
}
