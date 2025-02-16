package bins

import (
	"encoding/json"
	"fmt"
	"homeWork/3-struct/file"
	"time"
)

// Переменная хранящая лист Bin
var BinList = []Bin{}

// Структура Bin со структурными тегами для JSON
type Bin struct {
	Id        string    `json:"id"`
	Private   bool      `json:"private"`
	CreatedAt time.Time `json:"createdAt"`
	Name      string    `json:"name"`
}

// Метод по преобразованию в JSON
func (b *Bin) ToBytes() ([]byte, error) {
	file, err := json.Marshal(b)
	if err != nil {
		return nil, err
	}
	return file, nil
}

// Метод по добавлению bin в JSON и метод по удалению(обновлению) JSON
func (b *Bin) AddBin() error {
	b.CreatedAt = time.Now()
	data, err := b.ToBytes()
	if err != nil {
		return fmt.Errorf("Ошибка преобразования %v", err)
	}
	return file.WriteFile(data, "data.json")

}

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
	return &Bin{
		Id:        params.Id,
		Private:   params.Private,
		CreatedAt: time.Now(),
		Name:      params.Name,
	}, nil
}
