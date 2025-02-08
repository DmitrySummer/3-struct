package file

import (
	"encoding/json"
	"fmt"
	"os"
)

type File interface {
	WriteFile([]byte, string)
}

// Функция записи файла и проверка является ли файл JSON
func WriteFile(content []byte, name string) error {
	if !json.Valid(content) {
		return fmt.Errorf("Некорректный JSON, ошибка записи")
	}
	var existingData []map[string]interface{}
	if data, err := os.ReadFile(name); err == nil && len(data) > 0 {
		json.Unmarshal(data, &existingData)
	}
	var newData map[string]interface{}
	if err := json.Unmarshal(content, &newData); err != nil {
		return fmt.Errorf("Ошибка декодирования JSON: %v", err)
	}

	existingData = append(existingData, newData)

	updatedData, err := json.MarshalIndent(existingData, "", "  ")
	if err != nil {
		return fmt.Errorf("Ошибка кодирования JSON: %v", err)
	}

	file, err := os.OpenFile(name, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	file.Truncate(0)
	_, err = file.Write(updatedData)
	if err != nil {
		return err
	}

	fmt.Println("Данные успешно добавлены в", name)
	return nil
}
