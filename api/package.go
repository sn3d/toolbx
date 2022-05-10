package api

type Package struct {
	Platform string `yaml:"platform"`
	Uri      string `yaml:"uri"`
	Cmd      string `yaml:"cmd",omitempty`
}
