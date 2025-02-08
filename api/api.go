package api

import (
	"fmt"
	"homeWork/3-struct/config"
)

func ReadApi(keyFile string) (string, error) {
	conf := config.Config{}

	key, err := conf.ReadKey(keyFile)
	if err != nil {
		return "", fmt.Errorf("Не удалось прочитать ключ api %v", err)
	}
	return key, nil
}
