package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Key string
}

func NewConfig() *Config {
	return &Config{}
}

func (conf *Config) LoadEnv(fileEnv string) error {
	err := godotenv.Load(fileEnv)
	if err != nil {
		log.Panic("не удалось прочитать файл .env", err)
	}
	conf.Key = os.Getenv("KEY")
	return nil
}

func (conf *Config) GetKey() string {
	return conf.Key
}
