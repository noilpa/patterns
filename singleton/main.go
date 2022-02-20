package main

import (
	"fmt"
	"sync"

	"patterns/singleton/singleton"
)

func main() {
	wg := new(sync.WaitGroup)
	wg.Add(1)
	go func() {
		defer wg.Done()
		s := singleton.New()
		s.Inc()
		s.Inc()
		s.Inc()
		s.Inc()
		fmt.Println("s", s.Cntr())
		fmt.Printf("&s=%p\n", s)
	}()

	s1 := singleton.New()
	fmt.Println("s1", s1.Cntr())
	s1.Inc()
	s1.Inc()
	fmt.Println("s1", s1.Cntr())

	s2 := singleton.New()
	fmt.Println("s2", s2.Cntr())
	s2.Inc()
	fmt.Println("s1", s1.Cntr())

	fmt.Printf("&s1=%p &s2=%p\n", s1, s2)
	wg.Wait()
}
