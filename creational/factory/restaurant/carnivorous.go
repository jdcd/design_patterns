package restaurant

type BeefKing struct{}

func (r *BeefKing) Buy() (*Menu, int) {
	return &Menu{Lunch: "Beef"}, 120
}
