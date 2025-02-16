package api

import (
	"encoding/json"
	"fmt"
	"homeWork/3-struct/bins"
	"homeWork/3-struct/storage"
)

// Создание Bin - POST
func PostBin(params *bins.BinParams) error {
	bin, err := bins.NewBin(params)
	if err != nil {
		return fmt.Errorf("Не удалось создать Bin: %v", err)
	}

	if bin.CreatedAt == "" {
		bin.CreatedAt = bins.GetCurrentTime()
	}

	if err := bin.AddBin("data.json"); err != nil {
		return fmt.Errorf("Ошибка сохранения Bin: %v", err)
	}

	fmt.Println("Bin успешно создан:", bin.Id)
	return nil
}

// Удаление Bin - DEL
func DelBin(binID string) error {
	return bins.DeleteBinByID("data.json", binID)
}

// Получение Bin - GET
func GetBin(fileName string) (*bins.Bin, error) {
	data, err := storage.ReadFile(fileName)
	if err != nil {
		return nil, fmt.Errorf("Ошибка чтения файла: %w", err)
	}

	var bin bins.Bin
	if err := json.Unmarshal(data, &bin); err != nil {
		return nil, fmt.Errorf("Ошибка парсинга JSON: %w", err)
	}

	return &bin, nil
}

// Обновление Bin - PUT
func PutBin(fileName, binID, newName string, newPrivate bool) (*bins.Bin, error) {
	data, err := storage.ReadFile(fileName)
	if err != nil {
		return nil, fmt.Errorf("Ошибка чтения файла: %w", err)
	}

	var binsList []bins.Bin
	if len(data) > 0 {
		if err := json.Unmarshal(data, &binsList); err != nil {
			return nil, fmt.Errorf("Ошибка парсинга JSON: %w", err)
		}
	}

	var updatedBin *bins.Bin
	for i := range binsList {
		if binsList[i].Id == binID {
			binsList[i].Name = newName
			binsList[i].Private = newPrivate
			binsList[i].CreatedAt = bins.GetCurrentTime()
			updatedBin = &binsList[i]
			break
		}
	}

	if updatedBin == nil {
		return nil, fmt.Errorf("Bin с таким ID %s не найден", binID)
	}

	if err := storage.WriteJSON(fileName, binsList); err != nil {
		return nil, fmt.Errorf("Ошибка обновления файла: %v", err)
	}

	fmt.Println("Bin успешно обновлен:", binID)
	return updatedBin, nil
}
