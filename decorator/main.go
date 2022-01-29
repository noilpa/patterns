package main

import (
	"fmt"

	"patterns/decorator/beverage"
	"patterns/decorator/topping"
)

// Here we have a Coffee point with "house blend" and "dark roast" coffee.
// And we have a two toppings: milk and mocha.
//
// We need to implement ticket window with drink description and cost.

func main() {
	var coffee beverage.Beverage
	coffee = beverage.NewDarkRoast()
	fmt.Println(coffee.GetDescription(), "|", coffee.Cost())

	var coffee1 beverage.Beverage
	coffee1 = beverage.NewHouseBlend()
	coffee1 = topping.NewMilk(coffee1)
	fmt.Println(coffee1.GetDescription(), "|", coffee1.Cost())

	var coffee2 beverage.Beverage
	coffee2 = beverage.NewDarkRoast()
	coffee2 = topping.NewMocha(coffee2)
	coffee2 = topping.NewMilk(coffee2)
	fmt.Println(coffee2.GetDescription(), "|", coffee2.Cost())
}
