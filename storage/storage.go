package storage

import (
	"encoding/json"
	"fmt"
	"os"
)

type FileStorage interface {
	ReadFile(string) ([]byte, error)
	SaveFile(string, map[string]interface{})
}

// Функция для прочтения файла
func ReadFile(name string) ([]byte, error) {
	data, err := os.ReadFile(name)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// Функция по сохранению файла, если файл уже существует, то добавит запись уже в имеющийся
func SaveFile(name string, newData map[string]interface{}) error {
	var existingData []map[string]interface{}
	file, err := os.ReadFile(name)
	if err == nil && len(file) > 0 {
		if err := json.Unmarshal(file, &existingData); err != nil {
			return fmt.Errorf("ошибка декодирования JSON: %v", err)
		}
	}
	existingData = append(existingData, newData)
	updatedData, err := json.Marshal(existingData)
	if err != nil {
		return fmt.Errorf("ошибка кодирования JSON: %v", err)
	}

	if err := os.WriteFile(name, updatedData, 0); err != nil {
		return fmt.Errorf("ошибка записи в файл: %v", err)
	}

	fmt.Println("Данные успешно добавлены в", name)
	return nil
}
