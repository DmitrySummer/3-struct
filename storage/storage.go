package storage

import (
	"encoding/json"
	"fmt"
	"os"
)

// Функция для чтения файла
func ReadFile(name string) ([]byte, error) {
	return os.ReadFile(name)
}

// Функция для записи файла JSON
func WriteJSON(fileName string, data interface{}) error {
	updatedData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return fmt.Errorf("Ошибка кодирования JSON: %v", err)
	}

	return os.WriteFile(fileName, updatedData, 0644)
}

// Функция для сохранения нового Bin или обновления существующего
func SaveFile(fileName string, newData map[string]interface{}) error {
	var existingData []map[string]interface{}

	file, err := os.ReadFile(fileName)
	if err == nil && len(file) > 0 {
		if file[0] == '{' {
			var singleObject map[string]interface{}
			if err := json.Unmarshal(file, &singleObject); err == nil {
				existingData = append(existingData, singleObject)
			} else {
				return fmt.Errorf("Ошибка декодирования JSON: %v", err)
			}
		} else {
			if err := json.Unmarshal(file, &existingData); err != nil {
				return fmt.Errorf("Ошибка декодирования JSON: %v", err)
			}
		}
	}

	found := false
	for i, item := range existingData {
		if id, ok := item["id"].(string); ok && id == newData["id"] {
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
		return fmt.Errorf("Ошибка кодирования JSON: %v", err)
	}

	if err := os.WriteFile(fileName, updatedData, 0644); err != nil {
		return fmt.Errorf("Ошибка записи в файл: %v", err)
	}

	fmt.Println("Данные успешно обновлены в", fileName)
	return nil
}
