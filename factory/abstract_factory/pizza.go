package abstract_factory

import "fmt"

type PizzaType string

const (
	Cheese    PizzaType = "cheese"
	Pepperoni PizzaType = "pepperoni"
	Veggie    PizzaType = "veggie"
)

type Pizza interface {
	Prepare()
	Bake()
	Cut()
	Box()
}

type cheesePizza struct {
	pif PizzaIngredientFactory
}

func (c *cheesePizza) Prepare() {
	c.pif.CreateSauce()
	c.pif.CreateCheese()
	return
}

func (c *cheesePizza) Bake() {
	fmt.Println("bake cheese pizza")
}

func (c *cheesePizza) Cut() {
	fmt.Println("cut cheese pizza")
}

func (c *cheesePizza) Box() {
	fmt.Println("box cheese pizza")
}

func NewCheesePizza(pif PizzaIngredientFactory) *cheesePizza {
	return &cheesePizza{pif: pif}
}

type pepperoniPizza struct {
	pif PizzaIngredientFactory
}

func (c *pepperoniPizza) Prepare() {
	c.pif.CreateSauce()
	c.pif.CreatePepperoni()
	return
}

func (c *pepperoniPizza) Bake() {
	fmt.Println("bake pepperoni pizza")
}

func (c *pepperoniPizza) Cut() {
	fmt.Println("cut pepperoni pizza")
}

func (c *pepperoniPizza) Box() {
	fmt.Println("box pepperoni pizza")
}

func NewPepperoniPizza(pif PizzaIngredientFactory) *pepperoniPizza {
	return &pepperoniPizza{pif: pif}
}

type veggiePizza struct {
	pif PizzaIngredientFactory
}

func (c *veggiePizza) Prepare() {
	c.pif.CreateSauce()
	c.pif.CreateVeggies()
	return
}

func (c *veggiePizza) Bake() {
	fmt.Println("bake veggie pizza")
}

func (c *veggiePizza) Cut() {
	fmt.Println("cut veggie pizza")
}

func (c *veggiePizza) Box() {
	fmt.Println("box veggie pizza")
}

func NewVeggiePizza(pif PizzaIngredientFactory) *veggiePizza {
	return &veggiePizza{pif: pif}
}
