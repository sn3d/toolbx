package config

type Configuration struct {
	// URL to Git repository with commands you want to sync with (e.g. https://gitlab.myorg.com/toolbx-commands.git
	CommandsRepository string `yaml:"repository",omitempty`

	// Git's branch that will be used for syncing command to local machine from remote Git
	CommandsBranch string `yaml:"branch",omitempty`

	// GitLab token used for cloning command repository from GitLab or getting tools hosted
	// in GitLab as artifacts
	Token string

	// Absolute path to data directory where are stored installed tools,
	// synced repo etc.
	DataDir string `yaml:"data",omitempty`

	// Here you can change 'toolbx' with your brand label. It's for whitelabeling purpose
	BrandLabel string
}
