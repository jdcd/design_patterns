package factory

import (
	"fmt"
	"github.com/jdcd9001/design_patterns_go/creational/factory/restaurant"
	"github.com/jdcd9001/design_patterns_go/creational/factory/restaurant/kitchen"
)

func GetRestaurant(preference kitchen.KindOfFood) (restaurant.IRestaurant, error) {
	switch preference {
	case kitchen.Carnivorous:
		return &restaurant.BeefKing{}, nil
	case kitchen.Vegetarian:
		return &restaurant.Vegy{}, nil
	default:
		return nil, fmt.Errorf("we don't offer that kind of food: %s", preference)
	}
}
