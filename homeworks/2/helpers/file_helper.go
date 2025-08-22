package helpers

import (
	"fmt"
	"github.com/Alex43063/training_go/homeworks/2/actions"
	"os"
)

func SaveToFile(inv *actions.Inventory, filename string) {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Ошибка при создании файла:", err)
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Printf("Ошибка при закрытие файла %v", err)
			return
		}
	}(file)
	inv.Save(file)
}

func LoadFromFile(filename string) actions.Inventory {
	newInv := actions.Inventory{}
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Ошибка при открытии файла:", err)
		return actions.Inventory{}
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("Ошибка при закрытие файла:", err)
		}
	}(file)
	newInv.Load(file)
	return newInv
}
