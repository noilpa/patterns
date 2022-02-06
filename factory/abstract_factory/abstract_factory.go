package abstract_factory

import "fmt"


type PizzaStore interface {
	CreatePizza(pt PizzaType) (Pizza, error)
}

type NyPizzaStore struct {
	pif PizzaIngredientFactory // NY pizza ingredients
}

func NewNyPizzaStore(pif PizzaIngredientFactory) *NyPizzaStore {
	return &NyPizzaStore{pif: pif}
}

func (c *NyPizzaStore) CreatePizza(pt PizzaType) (Pizza, error) {
	var pizza Pizza
	switch pt {
	case Cheese:
		pizza = NewCheesePizza(c.pif)
	case Pepperoni:
		pizza = NewPepperoniPizza(c.pif)
	case Veggie:
		pizza = NewVeggiePizza(c.pif)
	default:
		return nil, fmt.Errorf("unknown pizza type: %s", pt)
	}

	pizza.Prepare()
	pizza.Bake()
	pizza.Box()
	pizza.Cut()

	return pizza, nil
}

type ChicagoPizzaStore struct {
	pif PizzaIngredientFactory // Chicago pizza ingredients
}

func NewChicagoPizzaStore(pif PizzaIngredientFactory) *ChicagoPizzaStore {
	return &ChicagoPizzaStore{pif: pif}
}

func (c *ChicagoPizzaStore) CreatePizza(pt PizzaType) (Pizza, error) {
	var pizza Pizza
	switch pt {
	case Cheese:
		pizza = NewCheesePizza(c.pif)
	case Pepperoni:
		pizza = NewPepperoniPizza(c.pif)
	case Veggie:
		pizza = NewVeggiePizza(c.pif)
	default:
		return nil, fmt.Errorf("unknown pizza type: %s", pt)
	}

	pizza.Prepare()
	pizza.Bake()
	pizza.Box()
	pizza.Cut()

	return pizza, nil
}