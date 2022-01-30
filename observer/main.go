package main

import "fmt"

func main() {
	ss := &ConcreteSubject{
		observers: make(map[Observer]struct{}),
	}
	co1 := &ConcreteObserver{Name: "Alpha"}
	co2 := &ConcreteObserver{Name: "Beta"}

	ss.SetState("A")

	ss.Register(co1)
	ss.SetState("B")

	ss.Register(co2)
	ss.SetState("C")

	ss.Remove(co1)
	ss.SetState("D")
}

// what is observed
type Subject interface {
	Register(o Observer)
	Remove(o Observer)
	notify()
}

type ConcreteSubject struct {
	observers map[Observer]struct{}
	state     string
}

func (c *ConcreteSubject) Register(o Observer) {
	if _, exist := c.observers[o]; exist {
		return
	}
	c.observers[o] = struct{}{}
}

func (c *ConcreteSubject) Remove(o Observer) {
	if _, exist := c.observers[o]; exist {
		delete(c.observers, o)
	}
}

func (c *ConcreteSubject) SetState(s string) {
	fmt.Printf("Subject change state to: %s\n", s)
	c.state = s
	c.notify()
}

func (c *ConcreteSubject) notify() {
	for o := range c.observers {
		o.Handle(c.state)
	}
}

// who are observing
type Observer interface {
	Handle(msg string)
}

type ConcreteObserver struct {
	Name string
}

func (c *ConcreteObserver) Handle(msg string) {
	fmt.Printf("%s handle message from subject: %s\n", c.Name, msg)
}
