package guns

import (
	"fmt"
	"io"
)

type Armor struct {
	Name    string
	Defense int
	Weight  float64
}

func (a *Armor) Use() string {
	return fmt.Sprintf(
		"Защита от атаки с использованием '%s', защита составляет %d",
		a.Name,
		a.Defense,
	)
}

func (a *Armor) GetName() string {
	return a.Name
}

func (a *Armor) GetWeight() float64 {
	return a.Weight
}

func (a *Armor) Serialize(writer io.Writer) {
	_, err := writer.Write([]byte(fmt.Sprintf("Armor|%s|%d|%.2f\n", a.Name, a.Defense, a.Weight)))
	if err != nil {
		fmt.Printf("Ошибка сериализации %v", err)
		return
	}
}

func (a *Armor) Deserialize(r io.Reader) {
	_, err := fmt.Fscanf(r, "Armor|%s|%d|%.2f\n", &a.Name, &a.Defense, &a.Weight)
	if err != nil {
		fmt.Printf("Ошибка десериализации %v", err)
		return
	}
}
