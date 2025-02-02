package storage

import "os"

// Функция для прочтения файла
func ReadFile(name string) ([]byte, error) {
	data, err := os.ReadFile(name)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// Функция по сохранению файла
func SaveFile(name string, content []byte) error {
	err := os.WriteFile(name, content, 0644)
	if err != nil {
		return err
	}
	return nil
}
