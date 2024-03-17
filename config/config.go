package config

import (
	"log"
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
)

var once sync.Once
var configInstance *Config

type (
	Config struct {
		Token   `yaml:"token"`
	}


	Token struct {
		TokenExp string `env-required:"true" yaml:"token_exp" env:"TOKEN_EXPIRATION"`
		TokenKey string `env-required:"true" yaml:"token_key" env:"TOKEN_KEY"`
	}
)

func GetConfig() *Config {
	if configInstance == nil {
		log.Println("Error: Config not loaded")
	}
	return configInstance
}

func InitConfig(p string) *Config {
	once.Do(func() {
		cfg := &Config{}

		err := cleanenv.ReadConfig(p, cfg)
		if err != nil {
			log.Fatal("Config error: ", err)
		}

		err = cleanenv.ReadEnv(cfg)
		if err != nil {
			log.Fatal("Config error: ", err)
		}
		configInstance = cfg
	})

	return configInstance
}
