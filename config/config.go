package config

import (
	"github.com/kelseyhightower/envconfig"
	"log"
)

type Config struct {
	MySQL struct {
		Host     string `envconfig:"MYSQL_HOST"`
		User     string `envconfig:"MYSQL_USERNAME"`
		Password string `envconfig:"MYSQL_ROOT_PASSWORD"`
		DBName   string `envconfig:"MYSQL_DATABASE"`
	}
}

func Init() Config {
	var cfg Config
	err := envconfig.Process("", &cfg)
	if err != nil {
		log.Fatalln(err)
	}

	return cfg
}
