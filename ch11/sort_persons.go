package main

import (
	"fmt"
	"sort"
)

type Person struct {
	firstName string
	lastName  string
}
type Persons []Person

func (p Persons) Len() int { return len(p) }
func (p Persons) Less(i, j int) bool {
	in := p[i].lastName + " " + p[i].firstName
	jn := p[j].lastName + " " + p[j].firstName
	return in < jn
}
func (p Persons) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
func main() {
	p1 := Person{"Barack", "Obama"}
	p2 := Person{"Donald", "Trump"}
	p3 := Person{"Joe", "Biden"}
	arrP := Persons{p1, p2, p3}
	fmt.Println("Before sorting:\n", arrP)
	sort.Sort(arrP)
	fmt.Println("After sorting:\n", arrP)
}
