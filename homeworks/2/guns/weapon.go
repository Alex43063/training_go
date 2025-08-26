package guns

import (
	"fmt"
	"io"
)

type Weapon struct {
	Name       string
	Damage     int
	Durability int
}

func (w *Weapon) Use() string {
	if !(w.Durability == 0) {
		w.Durability -= 1
	}
	return fmt.Sprintf(
		"Используется оружие '%s' c уроном %d, прочность оружия после атаки = %d",
		w.Name,
		w.Damage,
		w.Durability,
	)
}

func (w *Weapon) GetName() string {
	return w.Name
}

func (w *Weapon) GetWeight() float64 {
	return 0.7 * float64(w.Durability+w.Damage)
}

func (w *Weapon) Serialize(writer io.Writer) {
	_, err := writer.Write([]byte(fmt.Sprintf("Weapon|%s|%d|%d\n", w.Name, w.Damage, w.Durability)))
	if err != nil {
		fmt.Printf("Ошибка сериализации %v", err)
		return
	}
}

func (w *Weapon) Deserialize(r io.Reader) {
	_, err := fmt.Fscanf(r, "Weapon|%s|%d|%d\n", &w.Name, &w.Damage, &w.Durability)
	if err != nil {
		fmt.Printf("Ошибка десериализации %v", err)
		return
	}
}
