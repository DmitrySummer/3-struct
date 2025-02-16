package bins

import (
	"encoding/json"
	"fmt"
	"homeWork/3-struct/storage"
	"time"
)

// Структура Bin
type Bin struct {
	Id        string `json:"id"`
	Private   bool   `json:"private"`
	CreatedAt string `json:"createdAt"`
	Name      string `json:"name"`
}

// Отдельная структура для параметров func NewBin
type BinParams struct {
	Id      string `json:"id"`
	Private bool   `json:"private"`
	Name    string `json:"name"`
}

// Функция по созданию нового Bin
func NewBin(params *BinParams) (*Bin, error) {
	if params.Id == "" || params.Name == "" {
		return nil, fmt.Errorf("Не введены значения для id и name")
	}

	bin := &Bin{
		Id:        params.Id,
		Private:   params.Private,
		CreatedAt: GetCurrentTime(),
		Name:      params.Name,
	}

	fmt.Println("Создан новый Bin:", bin)

	return bin, nil
}

// Преобразование в JSON
func (b *Bin) ToBytes() ([]byte, error) {
	return json.Marshal(b)
}

// Добавление Bin в JSON
func (b *Bin) AddBin(fileName string) error {
	data, err := json.MarshalIndent(b, "", "  ")
	if err != nil {
		return fmt.Errorf("Ошибка преобразования в JSON: %v", err)
	}
	var binMap map[string]interface{}
	if err = DecodingJson(data, &binMap); err != nil {
		return fmt.Errorf("Ошибка декодирования JSON: %v", err)
	}

	if err := storage.SaveFile(fileName, binMap); err != nil {
		return fmt.Errorf("Ошибка сохранения Bin: %v", err)
	}

	fmt.Println("Bin успешно сохранен в файл:", b.Id)
	return nil
}

// Удаление Bin по ID
func DeleteBinByID(fileName, binID string) error {
	data, err := storage.ReadFile(fileName)
	if err != nil {
		return fmt.Errorf("Ошибка чтения файла: %v", err)
	}

	if len(data) == 0 {
		return fmt.Errorf("Файл пустой, нечего удалять")
	}

	var binsList []map[string]interface{}
	if err = DecodingJson(data, &binsList); err != nil {
		return fmt.Errorf("Ошибка декодирования JSON: %v", err)
	}

	newBins := []map[string]interface{}{}
	found := false

	for _, bin := range binsList {
		if id, ok := bin["id"].(string); ok && id != binID {
			newBins = append(newBins, bin)
		} else if ok && id == binID {
			found = true
		}
	}

	if !found {
		return fmt.Errorf("Bin с ID %s не найден", binID)
	}

	if err := storage.WriteJSON(fileName, newBins); err != nil {
		return fmt.Errorf("Ошибка обновления файла: %v", err)
	}

	fmt.Println("Bin успешно удален:", binID)
	return nil
}

// Функуция по декодированию из JSON
func DecodingJson[T any](d []byte, b *T) error {
	if err := json.Unmarshal(d, b); err != nil {
		return fmt.Errorf("Ошибка декодирования JSON: %v", err)
	}
	return nil
}

// Функция по установке времени
func GetCurrentTime() string {
	return time.Now().Format(time.RFC3339)
}
