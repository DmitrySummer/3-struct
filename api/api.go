package api

import (
	"fmt"
	"homeWork/3-struct/config"
)

func ReadApi(conf *config.Config) string {
	if conf.Key == "" {
		panic(fmt.Errorf("Не удалось прочитать API файл"))
	}
	return conf.Key
}
