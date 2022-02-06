package abstract_factory

import "fmt"

type PizzaIngredientFactory interface {
	CreateSauce()
	CreateCheese()
	CreateVeggies()
	CreatePepperoni()
}

type NyPizzaIngredientFactory struct{}

func (n *NyPizzaIngredientFactory) CreateSauce() {
	fmt.Println("NY sauce")
}

func (n *NyPizzaIngredientFactory) CreateCheese() {
	fmt.Println("NY cheese")
}

func (n *NyPizzaIngredientFactory) CreateVeggies() {
	fmt.Println("NY veggies")
}

func (n *NyPizzaIngredientFactory) CreatePepperoni() {
	fmt.Println("NY pepperoni")
}

type ChicagoPizzaIngredientFactory struct{}

func (c *ChicagoPizzaIngredientFactory) CreateSauce() {
	fmt.Println("chicago sauce")
}

func (c *ChicagoPizzaIngredientFactory) CreateCheese() {
	fmt.Println("chicago cheese")
}

func (c *ChicagoPizzaIngredientFactory) CreateVeggies() {
	fmt.Println("chicago veggies")
}

func (c *ChicagoPizzaIngredientFactory) CreatePepperoni() {
	fmt.Println("chicago pepperoni")
}
