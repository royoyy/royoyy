package main

import (
	"fmt"
	"sort"
)

type Person struct {
	firstname string
	lastname  string
}
type Persons []Person

func (p Persons) Len() int { return len(p) }
func (p Persons) Less(i, j int) bool {
	in := p[i].lastname + " " + p[i].firstname
	jn := p[j].lastname + " " + p[j].firstname
	return in < jn
}
func (p Persons) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
func main() {
	p1 := Person{"Barack", "Obama"}
	p2 := Person{"Donald", "Trump"}
	p3 := Person{"Joe", "Biden"}
	ps := Persons{p1, p2, p3}
	sort.Sort(ps)
	fmt.Println(ps)
}
