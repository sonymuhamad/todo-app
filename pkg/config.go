package pkg

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/subosito/gotenv"
)

type EnvConfig struct {
	ServerPort  int    `envconfig:"SERVER_PORT"`
	Env         string `envconfig:"ENV"`
	DatabaseUrl string `envconfig:"DATABASE_URL"`
}

func LoadEnvConfig() EnvConfig {
	_ = gotenv.Load()

	var cfg EnvConfig

	err := envconfig.Process("", &cfg)
	if err != nil {
		panic(err)
	}

	return cfg
}
