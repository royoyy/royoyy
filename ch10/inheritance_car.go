package main

import "fmt"

type Engine interface {
	Start()
	Stop()
}
type Car struct {
	wheelCount int
	Engine
}

func (c Car) Start() {
	fmt.Println("Started")
}
func (c Car) Stop() {
	fmt.Println("Stopped")
}
func (c Car) getWheel() int {
	return c.wheelCount
}

type Mercedes struct {
	Car
}

func (m Mercedes) sayHi() {
	fmt.Println("Hi, Mercedes")
}
func main() {
	m := &Mercedes{Car{4, nil}}
	m.Start()
	m.Stop()
	fmt.Println("wheelCount:", m.getWheel())
	m.sayHi()
}
