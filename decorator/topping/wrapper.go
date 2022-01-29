package topping

import "patterns/decorator/beverage"

type milk struct {
	b beverage.Beverage
}

func NewMilk(b beverage.Beverage) *milk {
	return &milk{
		b: b,
	}
}

func (m *milk) Cost() int {
	return m.b.Cost() + 20
}

func (m *milk) GetDescription() string {
	return m.b.GetDescription() + ", " + "Milk"
}

type mocha struct {
	b beverage.Beverage
}

func NewMocha(b beverage.Beverage) *mocha {
	return &mocha{
		b: b,
	}
}

func (m *mocha) Cost() int {
	return m.b.Cost() + 30
}

func (m *mocha) GetDescription() string {
	return m.b.GetDescription() + ", " + "Mocha"
}
