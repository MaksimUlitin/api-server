package store

//Config...
type Config struct {
	DataBaseURL string `toml:"data_base_url"`
}

//NewConfig...
func NewConfig() *Config {
	return &Config{}
}
