package api

type Metadata struct {
	Version     string    `yaml:"version"`
	Description string    `yaml:"description",omitempty`
	Usage       string    `yaml:"usage",omitempty`
	Packages    []Package `yaml:"packages"`
}
