package actions

import (
	"bufio"
	"fmt"
	gun "github.com/Alex43063/training_go/homeworks/2/guns"
	"io"
	"strconv"
	"strings"
)

type Item interface {
	GetName() string
	GetWeight() float64
}

func DescribeItem(i Item) string {
	if i == nil {
		return "Предмет отсутствует"
	}
	return fmt.Sprintf("%s (вес: %.1f)", i.GetName(), i.GetWeight())
}

func Filter[T any](items []T, predicate func(T) bool) []T {
	result := []T{}
	for _, item := range items {
		if predicate(item) {
			result = append(result, item)
		}
	}
	return result
}

func Map[T any, R any](items []T, transform func(T) R) []R {
	result := []R{}
	for _, item := range items {
		result = append(result, transform(item))
	}
	return result
}

func Find[T any](items []T, condition func(T) bool) (T, bool) {
	var zero T
	for _, item := range items {
		if condition(item) {
			return item, true
		}
	}
	return zero, false
}

type Inventory struct {
	Items []Item
}

func (inv *Inventory) AddItem(item Item) {
	inv.Items = append(inv.Items, item)
}

func (inv *Inventory) GetWeapons() []*gun.Weapon {
	filteredWeapon := Filter(inv.Items, func(item Item) bool {
		_, ok := item.(*gun.Weapon)
		return ok
	})
	return Map(filteredWeapon, func(i Item) *gun.Weapon {
		return i.(*gun.Weapon)
	})
}

func (inv *Inventory) GetBrokenItems() []Item {
	return Filter(inv.Items, func(item Item) bool {
		switch defItem := item.(type) {
		case *gun.Weapon:
			return defItem.Durability <= 0
		case *gun.Potion:
			return defItem.Charges <= 0
		}
		return false
	})
}

func (inv *Inventory) GetItemNames() []string {
	return Map(inv.Items, func(item Item) string {
		return item.GetName()
	})
}

func (inv *Inventory) FindItemByName(name string) (Item, bool) {
	return Find(inv.Items, func(item Item) bool {
		return item.GetName() == name
	})
}

func (inv *Inventory) Save(w io.Writer) {
	for _, item := range inv.Items {
		if storable, ok := item.(Storable); ok {
			storable.Serialize(w)
		}
	}
}

func (inv *Inventory) Load(r io.Reader) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		splitedLine := strings.Split(line, "|")
		switch splitedLine[0] {
		case "Weapon":
			damage, err := strconv.Atoi(splitedLine[2])
			if err != nil {
				fmt.Printf("Некорректное значение 'Damage'")
				return
			}
			durability, err := strconv.Atoi(splitedLine[3])
			if err != nil {
				fmt.Printf("Некорректное значение 'Damage'")
				return
			}
			weapon := &gun.Weapon{
				Name:       splitedLine[1],
				Damage:     damage,
				Durability: durability,
			}
			inv.AddItem(weapon)
		case "Armor":
			defense, err := strconv.Atoi(splitedLine[2])
			if err != nil {
				fmt.Printf("Некорректное значение 'Defense'")
				return
			}
			weight, err := strconv.ParseFloat(splitedLine[3], 64)
			if err != nil {
				fmt.Printf("Некорректное значение 'Weight'")
				return
			}
			armor := &gun.Armor{
				Name:    splitedLine[1],
				Defense: defense,
				Weight:  weight,
			}
			inv.AddItem(armor)
		}
	}
}

type Storable interface {
	Serialize(w io.Writer)
	Deserialize(r io.Reader)
}
