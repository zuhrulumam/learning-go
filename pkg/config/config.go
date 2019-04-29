package config

import (
	"github.com/caarlos0/env"
)

type config struct {
	ListenAddr string `env:"LEARNING_GO_LISTEN_ADDR"`
	DbURL      string `env:"LEARNING_GO_DB_URL"`
}

var conf config

func ParseEnv() error {
	err := env.Parse(&conf)

	return err
}

func ListenAddr() string {
	return conf.ListenAddr
}

func DbURL() string {
	return conf.DbURL
}
