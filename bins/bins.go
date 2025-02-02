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
	CreatedAt time.Time `json:"CreatedAt"`
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

// Метод по добавлению bin в JSON
func (b *Bin) AddBin(bin Bin) {
	b.CreatedAt = time.Now()
	data, err := b.ToBytes()
	if err != nil {
		fmt.Println("Ошибка преобразования")
	}
	file.WriteFile(data, "data.json")
}

// Функция по созданию нового Bin
func NewBin(id string, private bool, name string) (*Bin, error) {
	return &Bin{
		Id:        id,
		Private:   private,
		CreatedAt: time.Now(),
		Name:      name,
	}, nil
}
