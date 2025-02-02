package main

import (
	"fmt"
	"homeWork/3-struct/bins"
	"homeWork/3-struct/file"
)

func main() {
	createNewFiles()
}

// Функция по созданию файла пользователем
func createNewFiles() {
	// Так как не было вводных данных к ДЗ о том как в итоге получать данные в нашу
	// программу и файл, для примера создал такой пробный вариант.
	id := promptData("Введите Ваш id")
	private := true
	name := promptData("Введите Ваше имя")

	newBin, err := bins.NewBin(id, private, name)
	if err != nil {
		fmt.Println("Ошибка создания Bin:", err)
		return
	}

	fileContent, err := newBin.ToBytes()
	if err != nil {
		fmt.Println("Ошибка перевода в JSON:", err)
		return
	}

	err = file.WriteFile(fileContent, "data.json")
	if err != nil {
		fmt.Println("Ошибка записи файла:", err)
	}
}

// Функция по обработке ввода пользователя и чтения введенных данных
func promptData(prompt string) string {
	fmt.Print(prompt + ": ")
	var res string
	fmt.Scan(&res)
	return res
}
