package api

import (
	"encoding/json"
	"fmt"
	"homeWork/3-struct/bins"
	"homeWork/3-struct/config"
	"homeWork/3-struct/storage"
)

func ReadApi(conf *config.Config) string {
	if conf.Key == "" {
		panic(fmt.Errorf("Не удалось прочитать API файл"))
	}
	return conf.Key
}

// Создание bin POST
func PostBin(params *bins.BinParams) error {
	bin, err := bins.NewBin(params)
	if err != nil {
		return fmt.Errorf("Не удалось создадть Bin %v", err)
	}
	if err := bin.AddBin(); err != nil {
		return fmt.Errorf("Нет данных для добавления в файл %v", err)
	}
	return nil
}

// Получение bin GET
func GetBin(fileName string) (*bins.Bin, error) {
	data, err := storage.ReadFile(fileName)
	if err != nil {
		return nil, fmt.Errorf("ошибка чтения файла: %w", err)
	}
	var bin bins.Bin
	if err := json.Unmarshal(data, &bin); err != nil {
		return nil, fmt.Errorf("ошибка парсинга JSON: %w", err)
	}

	return &bin, nil
}

// Обновление bin PUT
func PutBin(fileName string) (file *bins.Bin, err error) {
	bin, err := GetBin(fileName)
	if err != nil {
		return nil, fmt.Errorf("Не удалось прочитать файл")
	}
	if err := bin.AddBin(); err != nil {
		return nil, fmt.Errorf("ошибка обновления: %v", err)
	}
	return bin, nil
}
