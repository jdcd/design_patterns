package factory

import (
	"github.com/jdcd9001/design_patterns_go/creational/factory/restaurant/kitchen"
	"testing"
)

func TestFoodFactory(t *testing.T) {
	tc := []struct {
		people []People
		exp    int
		name   string
	}{
		{[]People{
			{id: 1, name: "lady", preference: kitchen.Carnivorous},
			{id: 2, name: "doge", preference: kitchen.Vegetarian},
			{id: 3, name: "pandora", preference: "fish"}},
			140, "standard"},
		{[]People{
			{id: 1, name: "lady", preference: "wooden"},
			{id: 2, name: "doge", preference: "japanese without rice"},
			{id: 3, name: "michi", preference: "whiskas"}},
			0, "strange people"},
	}

	c := 0
	for _, i := range tc {
		c = MakeEventWithFood(i.people)
		if c != i.exp {
			t.Errorf("Error in test: %s, got: %d, exp: %d", i.name, c, i.exp)
		}
	}

}
