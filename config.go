package gwpro

type GwConfiguration struct {
	Store struct {
		Backend string `json:"backend" yaml:"backend" toml:"backend"`
		Cache   string `json:"cache" yaml:"cache" toml:"cache"`
	} `json:"store" yaml:"store" toml:"store"`
}
