package restaurant

type Vegy struct{}

func (r *Vegy) Buy() (*Menu, int) {
	return &Menu{Lunch: "Tofu"}, 20
}
