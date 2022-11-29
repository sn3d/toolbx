package command

// Tool's package that can be installed
type Package struct {
	Platform string `yaml:"platform"`
	Uri      string `yaml:"uri"`
	Cmd      string `yaml:"cmd",omitempty`
}
