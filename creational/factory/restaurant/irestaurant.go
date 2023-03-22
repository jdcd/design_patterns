package restaurant

type IRestaurant interface {
	Buy() (lunch *Menu, price int)
}
