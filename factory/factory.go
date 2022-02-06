package main

import (
	"fmt"

	af "patterns/factory/abstract_factory"
	fm "patterns/factory/factory_method"
)

func main() {
	vps := new(fm.VipPizzaStore)
	if _, err := vps.CreatePizza(fm.Cheese); err != nil {
		fmt.Println(err)
	}
	if _, err := vps.CreatePizza(fm.Pepperoni); err != nil {
		fmt.Println(err)
	}
	if _, err := vps.CreatePizza("vegggie"); err != nil {
		fmt.Println(err)
	}

	fmt.Println()
	fmt.Println()

	nps := af.NewNyPizzaStore(new(af.NyPizzaIngredientFactory))
	cps := af.NewChicagoPizzaStore(new(af.ChicagoPizzaIngredientFactory))

	if _, err :=nps.CreatePizza(af.Cheese); err != nil {
		fmt.Println(err)
	}
	if _, err := cps.CreatePizza(af.Cheese); err != nil {
		fmt.Println(err)
	}
}
