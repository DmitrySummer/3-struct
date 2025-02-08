package config

import (
	"fmt"
	"os"
)

type Config struct {
	Key string
}

func (conf *Config) ReadKey(keyFile string) (string, error) {
	data, err := os.ReadFile(keyFile)
	if err != nil {
		return "", fmt.Errorf("Не удалось прочитать файл .env %s: %v", keyFile, err)
	}
	conf.Key = string(data)
	return conf.Key, nil
}
