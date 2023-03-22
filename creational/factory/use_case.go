package factory

import (
	"fmt"
	"github.com/jdcd9001/design_patterns_go/creational/factory/factory"
	"github.com/jdcd9001/design_patterns_go/creational/factory/restaurant"
	"github.com/jdcd9001/design_patterns_go/creational/factory/restaurant/kitchen"
)

type People struct {
	id         int
	name       string
	preference kitchen.KindOfFood
}

func MakeEventWithFood(people []People) (check int) {
	var r restaurant.IRestaurant
	var err error
	for _, p := range people {
		r, err = factory.GetRestaurant(p.preference)
		if err != nil {
			fmt.Printf("the person: %s with id : %d, does not eat, reason: %s\n", p.name, p.id, err)
			continue
		}
		menu, price := r.Buy()
		check += price
		fmt.Printf("the person: %s with id : %d, ate:%s for %d USD\n", p.name, p.id, *menu, price)
	}

	return check
}
