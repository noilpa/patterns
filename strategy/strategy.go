package main

import "fmt"

func main() {
	md := New(new(flyWithWings), new(quack))
	md.display()
	md.fly()
	md.quack()

	// change behavior at runtime
	md.f = new(flyNoWay)
	md.q = new(quackMute)
	md.display()
	md.fly()
	md.quack()
}

type flyBehavior interface {
	fly()
}

type flyWithWings struct{}

func (f *flyWithWings) fly() {
	fmt.Println("fly with wings")
}

type flyNoWay struct {}

func (f *flyNoWay) fly() {
	fmt.Println("fly no way")
}

type quackBehavior interface {
	quack()
}

type quack struct {}

func (q *quack) quack() {
	fmt.Println("QUACK")
}

type quackMute struct {}

func (q *quackMute) quack() {
	fmt.Println("...(quiet)...")
}

type duck interface {
	fly()
	quack()
	display()
}

type myDuck struct {
	f flyBehavior
	q quackBehavior
}

func New(f flyBehavior, q quackBehavior) *myDuck {
	return &myDuck{
		f: f,
		q: q,
	}
}

func (m *myDuck) fly() {
	m.f.fly()
}

func (m *myDuck) quack() {
	m.q.quack()
}

func (m *myDuck) display() {
	fmt.Println("my duck!!!")
}

