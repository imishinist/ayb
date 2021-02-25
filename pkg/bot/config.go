package bot

import "github.com/kelseyhightower/envconfig"

type Config struct {
	Env string `envconfig:"ENV" default:"dev"`
}

func loadConfig() *Config {
	var conf Config
	if err := envconfig.Process("ayb", &conf); err != nil {
		panic(err)
	}
	return &conf
}
