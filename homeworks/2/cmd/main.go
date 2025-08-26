package main

import (
	"fmt"
	"github.com/Alex43063/training_go/homeworks/2/actions"
	guns "github.com/Alex43063/training_go/homeworks/2/guns"
	"github.com/Alex43063/training_go/homeworks/2/helpers"
)

func main() {
	// TODO: Создайте инвентарь и добавьте:
	// TODO: - Оружие: "Меч" (урон 10, прочность 5)
	// TODO: - Броню: "Щит" (защита 5, вес 4.5)
	// TODO: - Зелье: "Лечебное" (+50 HP, 3 заряда)
	// TODO: - Оружие: "Сломанный лук" (урон 5, прочность 0)
	inv := actions.Inventory{}

	inv.AddItem(&guns.Weapon{Name: "Меч", Damage: 10, Durability: 5})
	inv.AddItem(&guns.Armor{Name: "Щит", Defense: 5, Weight: 4.5})
	inv.AddItem(&guns.Potion{Name: "Лечебное", Effect: "+50 HP", Charges: 3})
	inv.AddItem(&guns.Weapon{Name: "Сломанный лук", Damage: 5, Durability: 0})

	// TODO: Реализуйте логику/вызовы:
	// TODO: 1. Use предмета с выводом в консоль
	for _, item := range inv.Items {
		switch defItem := item.(type) {
		case *guns.Weapon:
			fmt.Println(defItem.Use())
		case *guns.Potion:
			fmt.Println(defItem.Use())
		case *guns.Armor:
			fmt.Println(defItem.Use())
		}
	}

	// TODO: 2. DescribeItem с предметом и с nil, так же с выводом в консоль
	for _, item := range inv.Items {
		fmt.Println(actions.DescribeItem(item))
	}
	fmt.Println(actions.DescribeItem(nil))
	// TODO: 3. Вывести в консоль результат вызова GetWeapons (должны вернуться только меч и лук)
	for _, item := range inv.GetWeapons() {
		fmt.Println(*item)
	}
	// TODO: 4. Вывести в консоль результат вызова GetBrokenItems (должен вернуть сломанный лук)

	for _, item := range inv.GetBrokenItems() {
		fmt.Println(item)
	}

	// TODO: 5. Вывести в консоль результат вызова GetItemNames (все названия)
	fmt.Println(inv.GetItemNames())
	// TODO: 6. Вывести в консоль результат вызова FindItemByName (поиск "Щит")
	fmt.Println(inv.FindItemByName("Щит"))

	// TODO: Бонус: сделайте сохранение инвентаря в файл и загрузку инвентаря из файла
	helpers.SaveToFile(&inv, "saved_inventory.txt")

	inventory := helpers.LoadFromFile("saved_inventory.txt")
	fmt.Println(inventory.GetItemNames())
}
