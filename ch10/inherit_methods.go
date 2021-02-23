package main

import "fmt"

type Base struct {
	id string
}

func (b *Base) Id() string {
	return b.id
}
func (b *Base) SetId(s string) {
	b.id = s
}

type Person struct {
	Base
	FirstName string
	LastName  string
}
type Employee struct {
	Person
	salary float32
}

func main() {
	e := &Employee{Person{Base{"a1"}, "donald", "trump"}, 1.5}
	fmt.Println(e)
	e.SetId("a2")
	fmt.Println("new id:", e.Id())
}
