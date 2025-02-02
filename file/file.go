package file

import (
	"encoding/json"
	"fmt"
	"os"
)

// Функция записи файла и проверка является ли файл JSON
func WriteFile(content []byte, name string) error {
	if !json.Valid(content) {
		return fmt.Errorf("Некорректный JSON, ошибка записи")
	}

	file, err := os.Create(name)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(content)
	if err != nil {
		return err
	}

	fmt.Println("Файл записался успешно")
	return nil
}
