package main

import (
	"fmt"
	"math"
)

type Shaper interface {
	Area() float32
}
type Shape struct{}

func (s Shape) Area() float32 {
	return -1
}

type Circle struct {
	radius float32
	Shape
}

func (c Circle) Area() float32 {
	return math.Pi * c.radius * c.radius
}
func main() {
	c := &Circle{1, Shape{}}
	shapers := []Shaper{c}
	for _, s := range shapers {
		fmt.Println(s.Area())
	}
}
