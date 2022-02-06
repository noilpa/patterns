package factory_method

import (
	"fmt"
)

type PizzaType string

const (
	Cheese    PizzaType = "cheese"
	Pepperoni PizzaType = "pepperoni"
	Veggie    PizzaType = "veggie"
)

type PizzaStore interface {
	CreatePizza(pt PizzaType) (Pizza, error)
}

type Pizza interface {
	Prepare()
	Bake()
	Cut()
	Box()
}

type VipPizzaStore struct{}

func (n *VipPizzaStore) CreatePizza(pt PizzaType) (Pizza, error) {
	var pizza Pizza
	switch pt {
	case Cheese:
		pizza = new(CheesePizza)
	case Pepperoni:
		pizza = new(PepperoniPizza)
	case Veggie:
		pizza = new(VeggiePizza)
	default:
		return nil, fmt.Errorf("unknown pizza type: %s", pt)
	}

	pizza.Prepare()
	pizza.Bake()
	pizza.Box()
	pizza.Cut()

	return pizza, nil
}

type CheesePizza struct{}

func (c *CheesePizza) Prepare() {
	fmt.Println("prepare cheese pizza")
}

func (c *CheesePizza) Bake() {
	fmt.Println("bake cheese pizza")
}

func (c *CheesePizza) Cut() {
	fmt.Println("cut cheese pizza")
}

func (c *CheesePizza) Box() {
	fmt.Println("box cheese pizza")
}

type PepperoniPizza struct{}

func (p *PepperoniPizza) Prepare() {
	fmt.Println("prepare pepperoni pizza")
}

func (p *PepperoniPizza) Bake() {
	fmt.Println("bake pepperoni pizza")
}

func (p *PepperoniPizza) Cut() {
	fmt.Println("cut pepperoni pizza")
}

func (p *PepperoniPizza) Box() {
	fmt.Println("box pepperoni pizza")
}

type VeggiePizza struct{}

func (v *VeggiePizza) Prepare() {
	fmt.Println("prepare veggie pizza")
}

func (v *VeggiePizza) Bake() {
	fmt.Println("bake veggie pizza")
}

func (v *VeggiePizza) Cut() {
	fmt.Println("cut veggie pizza")
}

func (v *VeggiePizza) Box() {
	fmt.Println("box veggie pizza")
}
