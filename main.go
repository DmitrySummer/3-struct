package main

import (
	"flag"
	"fmt"
	"homeWork/3-struct/api"
	"homeWork/3-struct/bins"
	"os"
)

func main() {
	createNewFiles()
}

// Функция по созданию файла пользователем
func createNewFiles() {
	createFlag := flag.NewFlagSet("create", flag.ExitOnError)
	createFile := flag.String("file", "", "Название файла")
	createNameBin := flag.String("name", "", "Название BIN файла")
	createId := flag.String("id", "", "Наименование ID")

	updateFlag := flag.NewFlagSet("update", flag.ExitOnError)
	updateFile := flag.String("file", "", "Название файла для обновления")
	updateId := flag.String("id", "", "ID для обновления")

	deleteFlag := flag.NewFlagSet("delate", flag.ExitOnError)
	deleteId := flag.String("id", "", "ID для удаления")

	getFlag := flag.NewFlagSet("get", flag.ExitOnError)
	getId := flag.String("id", "", "ID для удаления")

	listFlag := flag.NewFlagSet("list", flag.ExitOnError)

	comands := map[string]func(){
		"create": func() {
			createFlag.Parse(os.Args[1:])
			if *createNameBin == "" || *createId == "" {
				fmt.Println("Ошибка: --name и --id обязательны для create")
			}
			params := &bins.BinParams{
				Id:      *createId,
				Private: true,
				Name:    *createNameBin,
			}
			if err := api.PostBin(params); err != nil {
				fmt.Println("Ошибка создания:", err)
			}
			fmt.Println("Bin и File успешно созданы:", params.Id)

		},
	}

	// id := promptData("Введите Ваш id")
	// private := true
	// name := promptData("Введите Ваше имя")

	// newBin, err := bins.NewBin(id, private, name)
	// if err != nil {
	// 	fmt.Println("Ошибка создания Bin:", err)
	// 	return
	// }

	// fileContent, err := newBin.ToBytes()
	// if err != nil {
	// 	fmt.Println("Ошибка перевода в JSON:", err)
	// 	return
	// }

	// err = file.WriteFile(fileContent, "data.json")
	// if err != nil {
	// 	fmt.Println("Ошибка записи файла:", err)
	// }
}

// Функция по обработке ввода пользователя и чтения введенных данных
func promptData(prompt string) string {
	fmt.Print(prompt + ": ")
	var res string
	fmt.Scan(&res)
	return res
}
