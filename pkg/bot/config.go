package bot

import "github.com/kelseyhightower/envconfig"

type Config struct {
	Env string `envconfig:"ENV" default:"dev"`

	AccessToken       string `envconfig:"ACCESS_TOKEN"`
	AccessTokenSecret string `envconfig:"ACCESS_TOKEN_SECRET"`
	ConsumerKey       string `envconfig:"CONSUMER_KEY"`
	ConsumerSecret    string `envconfig:"CONSUMER_SECRET"`
}

func loadConfig() *Config {
	var conf Config
	if err := envconfig.Process("ayb", &conf); err != nil {
		panic(err)
	}
	return &conf
}
