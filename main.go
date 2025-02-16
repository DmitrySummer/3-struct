package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"homeWork/3-struct/api"
	"homeWork/3-struct/bins"
	"homeWork/3-struct/storage"
)

func main() {
	flagCommands()
}

// Функция для обработки поступающих команд по флагам
func flagCommands() {
	createFlag := flag.NewFlagSet("create", flag.ExitOnError)
	createNameBin := createFlag.String("name", "", "Название BIN файла")
	createId := createFlag.String("id", "", "ID BIN")

	updateFlag := flag.NewFlagSet("update", flag.ExitOnError)
	updateFile := updateFlag.String("file", "", "Файл для обновления")
	updateId := updateFlag.String("id", "", "Id для обновления")

	deleteFlag := flag.NewFlagSet("delete", flag.ExitOnError)
	deleteId := deleteFlag.String("id", "", "Id для удаления")

	getFlag := flag.NewFlagSet("get", flag.ExitOnError)
	getId := getFlag.String("id", "", "ID для получения")

	listFlag := flag.NewFlagSet("list", flag.ExitOnError)

	commands := map[string]func(){
		"create": func() {
			createFlag.Parse(os.Args[2:])
			if *createNameBin == "" || *createId == "" {
				fmt.Println("Ошибка ввода значений для флагов. Для флага create обязательно нужно заполнить значения --name и --id")
				os.Exit(1)
			}

			params := &bins.BinParams{
				Id:      *createId,
				Private: true,
				Name:    *createNameBin,
			}

			if err := api.PostBin(params); err != nil {
				fmt.Println("Ошибка создания Bin:", err)
				os.Exit(1)
			}

			fmt.Println("Ваш Bin создан:", *createId)
		},
		"update": func() {
			updateFlag.Parse(os.Args[2:])
			if *updateFile == "" || *updateId == "" {
				fmt.Println("Ошибка ввода значений для флагов. Для флага update обязательно нужно заполнить значения --file и --id")
				os.Exit(1)
			}

			fmt.Print("Пожалуйста введите новое имя BIN: ")
			var newName string
			fmt.Scanln(&newName)

			fmt.Print("Сделать BIN приватным? (true/false): ")
			var newPrivate bool
			fmt.Scanln(&newPrivate)

			if _, err := api.PutBin(*updateFile, *updateId, newName, newPrivate); err != nil {
				fmt.Println("Ошибка обновления BIN:", err)
				os.Exit(1)
			}
			fmt.Println("BIN обновлен:", *updateId)
		},
		"delete": func() {
			deleteFlag.Parse(os.Args[2:])
			if *deleteId == "" {
				fmt.Println("Ошибка ввода значений для флагов. Для флага delete обязательно нужно заполнить значение --id")
				os.Exit(1)
			}

			if err := api.DelBin(*deleteId); err != nil {
				fmt.Println("Ошибка при попытке удаления BIN:", err)
				os.Exit(1)
			}

			fmt.Println("BIN удален:", *deleteId)
		},
		"get": func() {
			getFlag.Parse(os.Args[2:])

			if *getId == "" {
				fmt.Println("Ошибка ввода значяений для флагов. Для флага get обязательно нужно заполнить значение --id")
				os.Exit(1)
			}

			data, err := storage.ReadFile("data.json")
			if err != nil {
				fmt.Println("Ошибка прочтения файла:", err)
				os.Exit(1)
			}

			var binsList []bins.Bin
			if err := json.Unmarshal(data, &binsList); err != nil {
				fmt.Println("Ошибка распаковки JSON:", err)
				os.Exit(1)
			}

			for _, bin := range binsList {
				if bin.Id == *getId {
					fmt.Printf("Bin найден: %+v\n", bin)
					return
				}
			}

			fmt.Println("Bin с ID", *getId, "не найден")
		},
		"list": func() {
			listFlag.Parse(os.Args[2:])

			data, err := storage.ReadFile("data.json")
			if err != nil {
				fmt.Println("Ошибка прочтения файла:", err)
				os.Exit(1)
			}

			var binsList []bins.Bin
			if err := json.Unmarshal(data, &binsList); err != nil {
				fmt.Println("Ошибка распаковки JSON:", err)
				os.Exit(1)
			}

			if len(binsList) == 0 {
				fmt.Println("Список пуст")
				return
			}

			fmt.Println("Список доступных BIN:")
			for _, bin := range binsList {
				fmt.Printf("- ID: %s, Name: %s, CreatedAt: %s\n", bin.Id, bin.Name, bin.CreatedAt)
			}
		},
	}

	if len(os.Args) < 2 {
		fmt.Println("Ошибка ввода значений для флагов. При вводе флагов обязательно нужно выбрать значение из create, update, delete, get, list")
		os.Exit(1)
	}

	if cmd, exists := commands[os.Args[1]]; exists {
		cmd()
	} else {
		fmt.Println("Ошибка ввода значений для флагов. При вводе флагов обязательно нужно выбрать значение из create, update, delete, get, list")
		os.Exit(1)
	}
}
