package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

type BrokenItemError struct {
	Item Item
	Err  error
}

func (e BrokenItemError) Error() string {
	return e.Err.Error()
}

// Кастомные ошибки
var (
	// TODO: Добавьте необходимые кастомные ошибки
	serializeError   = errors.New(fmt.Sprintf("failed to serialize object"))
	deserializeError = errors.New(fmt.Sprintf("failed to deserialize string"))
	brokenItemError  = errors.New(fmt.Sprintf("item is broken"))
)

type Item interface {
	Use() (string, error)
	GetName() string
	GetWeight() float64
}

type Storable interface {
	Serialize(w io.Writer) error
	Deserialize(r io.Reader) error
}

type Weapon struct {
	Name       string
	Damage     int
	Durability int
}

func (w *Weapon) Use() (string, error) {
	brokenError := BrokenItemError{
		Item: w,
		Err:  brokenItemError,
	}
	if w.Durability <= 0 {
		return "", brokenError
	}

	w.Durability--

	return fmt.Sprintf("Атаковали %s (%d урона)", w.Name, w.Damage), nil
}

func (w *Weapon) GetName() string {
	return w.Name
}

func (w *Weapon) GetWeight() float64 {
	return 2.5
}

func (w *Weapon) Serialize(wr io.Writer) error {
	_, err := fmt.Fprintf(wr, "Weapon|%s|%d|%d", w.Name, w.Damage, w.Durability)
	if err != nil {
		return fmt.Errorf("serialize erorr: %s", serializeError.Error())
	}
	return nil
}

func (w *Weapon) Deserialize(r io.Reader) error {
	data, readErr := io.ReadAll(r)
	if readErr != nil {
		return fmt.Errorf("read erorr: %s", readErr.Error())
	}
	parts := strings.Split(string(data), "|")

	w.Name = parts[1]
	var parseError error
	w.Damage, parseError = strconv.Atoi(parts[2])
	if parseError != nil {
		return fmt.Errorf("parse erorr: %s, parsing string '%s' error %s", deserializeError, parts[2], parseError.Error())
	}
	w.Durability, parseError = strconv.Atoi(parts[3])
	if parseError != nil {
		return fmt.Errorf("parse erorr: %s, parsing string '%s' error %s", deserializeError, parts[3], parseError.Error())
	}
	return nil
}

type Armor struct {
	Name    string
	Defense int
	Weight  float64
}

func (a *Armor) Use() (string, error) {
	return fmt.Sprintf("Надели %s (+%d защиты)", a.Name, a.Defense), nil
}

func (a *Armor) GetName() string {
	return a.Name
}

func (a *Armor) GetWeight() float64 {
	return a.Weight
}

func (a *Armor) Serialize(wr io.Writer) error {
	_, err := fmt.Fprintf(wr, "Armor|%s|%d|%f", a.Name, a.Defense, a.Weight)
	if err != nil {
		return fmt.Errorf("serialize erorr: %s", serializeError.Error())
	}
	return nil
}

func (a *Armor) Deserialize(r io.Reader) error {
	data, readErr := io.ReadAll(r)
	if readErr != nil {
		return fmt.Errorf("read erorr: %s", readErr.Error())
	}
	parts := strings.Split(string(data), "|")

	a.Name = parts[1]
	var parseError error
	a.Defense, parseError = strconv.Atoi(parts[2])
	if parseError != nil {
		return fmt.Errorf("parse erorr: %s, parsing string '%s' error %s", deserializeError, parts[2], parseError.Error())
	}
	a.Weight, parseError = strconv.ParseFloat(parts[3], 64)
	if parseError != nil {
		return fmt.Errorf("parse erorr: %s, parsing string '%s' error %s", deserializeError, parts[3], parseError.Error())
	}
	return nil
}

type Potion struct {
	Name    string
	Effect  string
	Charges int
}

func (p *Potion) Use() (string, error) {
	brokenError := BrokenItemError{
		Item: p,
		Err:  brokenItemError,
	}
	if p.Charges <= 0 {
		return "", brokenError
	}

	p.Charges--

	return fmt.Sprintf("Использовали %s (%s)", p.Name, p.Effect), nil
}

func (p *Potion) GetName() string {
	return p.Name
}

func (p *Potion) GetWeight() float64 {
	return 0.5
}

func DescribeItem(i Item) (string, error) {
	if i == nil {
		return "", fmt.Errorf("item %v is missing", i)
	}

	return fmt.Sprintf("%s (вес: %.1f)", i.GetName(), i.GetWeight()), nil
}

func Filter[T any](items []T, predicate func(T) bool) []T {
	var result []T

	for _, item := range items {
		if predicate(item) {
			result = append(result, item)
		}
	}

	return result
}

func Map[T any, R any](items []T, transform func(T) R) []R {
	result := make([]R, len(items))

	for i, item := range items {
		result[i] = transform(item)
	}

	return result
}

func Find[T any](items []T, condition func(T) bool) (T, bool) {
	for _, item := range items {
		if condition(item) {
			return item, true
		}
	}

	var zero T

	return zero, false
}

type Inventory struct {
	Items []Item
}

func (inv *Inventory) AddItem(item Item) error {
	if item == nil {
		return fmt.Errorf("item is nill")
	}
	inv.Items = append(inv.Items, item)
	return nil
}

func (inv *Inventory) GetWeapons() []*Weapon {
	weapons := Filter(inv.Items, func(item Item) bool {
		_, ok := item.(*Weapon)
		return ok
	})

	return Map(weapons, func(item Item) *Weapon {
		return item.(*Weapon)
	})
}

func (inv *Inventory) GetBrokenItems() []Item {
	return Filter(inv.Items, func(item Item) bool {
		switch v := item.(type) {
		case *Weapon:
			return v.Durability <= 0
		case *Potion:
			return v.Charges <= 0
		default:
			return false
		}
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

func (inv *Inventory) Save(w io.Writer) error {
	for _, item := range inv.Items {
		if storable, ok := item.(Storable); ok {
			errSer := storable.Serialize(w)
			if errSer != nil {
				return errSer
			}

			_, err := fmt.Fprintln(w)
			if err != nil {
				return fmt.Errorf("save error %w", err)
			}
		}
	}
	return nil
}

func (inv *Inventory) Load(r io.Reader) error {
	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, "Weapon") {
			var w Weapon

			r := strings.NewReader(line)

			errDes := w.Deserialize(r)
			if errDes != nil {
				return errDes
			}

			err := inv.AddItem(&w)
			if err != nil {
				return err
			}
		} else if strings.HasPrefix(line, "Armor") {
			var a Armor

			r := strings.NewReader(line)

			errDes := a.Deserialize(r)
			if errDes != nil {
				return errDes
			}

			err := inv.AddItem(&a)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func SafeUse(item Item) (result string, err error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recovered in SafeUse", r)
		}
	}()
	if item.GetName() == "Ящик Пандоры" {
		panic("Don't use this item")
	}

	return item.Use()
}

func main() {
	// TODO: Реализуйте логику/вызовы:
	// TODO: 1. Обработку ошибок везде
	// TODO: 2. Use предмета до потери прочности и обработку ошибки при потере прочности
	// TODO: 3. DescribeItem с предметом и с nil
	// TODO: 4. Обработку ошибок сохранения/загрузки в файл
	// TODO: 5. Обработку паники для "Ящика Пандоры"
	inv := Inventory{}

	sword := &Weapon{Name: "Меч", Damage: 10, Durability: 5}
	healthPotion := &Potion{Name: "Лечебное", Effect: "+50 HP", Charges: 3}
	pandoraBox := &Weapon{Name: "Ящик Пандоры", Damage: math.MaxInt, Durability: math.MaxInt}

	addSwordErr := inv.AddItem(sword)
	if addSwordErr != nil {
		return
	}
	addPotionErr := inv.AddItem(healthPotion)
	if addPotionErr != nil {
		return
	}
	addPandoraErr := inv.AddItem(pandoraBox)
	if addPandoraErr != nil {
		return
	}
	err := inv.AddItem(nil)
	if err != nil {
		fmt.Println(err)
	}

	_, usePandoraErr := SafeUse(pandoraBox)
	if usePandoraErr != nil {
		return
	}

	fmt.Println(DescribeItem(sword))
	fmt.Println(DescribeItem(nil))

	fmt.Println("\nСохраняем в файл")

	file, _ := os.OpenFile("homework_solved.txt", os.O_RDWR|os.O_CREATE, 0644)

	saveErr := inv.Save(file)
	if saveErr != nil {
		return
	}

	fmt.Println("Ломаем файл")

	fmt.Fprintf(file, "Weapon||")

	fmt.Println("Загружаем из файла")
	inv = Inventory{}

	file, _ = os.Open("homework_solved.txt")

	loadErr := inv.Load(file)
	if loadErr != nil {
		fmt.Println(loadErr)
	}

	names := inv.GetItemNames()

	fmt.Println("\nИмена предметов:", names)

	for _, item := range inv.Items {
		describeItem, _ := DescribeItem(item)
		fmt.Println("-", describeItem)
	}
}
