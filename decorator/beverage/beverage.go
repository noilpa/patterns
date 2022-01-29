package beverage

type Beverage interface {
	Cost() int
	GetDescription() string
}

type houseBlend struct{}

func NewHouseBlend() *houseBlend {
	return &houseBlend{}
}

func (h *houseBlend) Cost() int {
	return 100
}

func (h *houseBlend) GetDescription() string {
	return "House Blend"
}

type darkRoast struct{}

func NewDarkRoast() *darkRoast {
	return &darkRoast{}
}

func (d *darkRoast) Cost() int {
	return 120
}

func (d *darkRoast) GetDescription() string {
	return "Dark Roast"
}
