package guns

import "fmt"

type Potion struct {
	Name    string
	Effect  string
	Charges int
}

func (p *Potion) Use() string {
	p.Charges -= 1
	return fmt.Sprintf(
		"Используется зелье '%s' c эфектом %s, осталось %d ед.",
		p.Name,
		p.Effect,
		p.Charges,
	)
}

func (p *Potion) GetName() string {
	return p.Name
}

func (p *Potion) GetWeight() float64 {
	return 0.1 * float64(p.Charges)
}
