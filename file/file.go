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
		return fmt.Errorf("некорректный JSON, ошибка записи")
	}

	var existingData []map[string]interface{}
	if data, err := os.ReadFile(name); err == nil && len(data) > 0 {
		if err := json.Unmarshal(data, &existingData); err != nil {
			return fmt.Errorf("ошибка чтения JSON: %v", err)
		}
	}

	var newData map[string]interface{}
	if err := json.Unmarshal(content, &newData); err != nil {
		return fmt.Errorf("ошибка декодирования JSON: %v", err)
	}

	id, exists := newData["id"]
	if !exists {
		return fmt.Errorf("JSON не содержит ключ 'id'")
	}

	found := false
	for i, item := range existingData {
		if item["id"] == id {
			existingData[i] = newData
			found = true
			break
		}
	}

	if !found {
		existingData = append(existingData, newData)
	}

	updatedData, err := json.MarshalIndent(existingData, "", "  ")
	if err != nil {
		return fmt.Errorf("ошибка кодирования JSON: %v", err)
	}

	file, err := os.OpenFile(name, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return fmt.Errorf("ошибка открытия файла: %v", err)
	}
	defer file.Close()

	_, err = file.Write(updatedData)
	if err != nil {
		return fmt.Errorf("ошибка записи файла: %v", err)
	}

	fmt.Println("Данные успешно обновлены в", name)
	return nil
}
