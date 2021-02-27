package main

import (
	"fmt"
	"math"
)

type Magnitude interface {
	Abs() float64
}

var m Magnitude

type Point struct {
	X, Y float64
}

func (p *Point) Scale(s float64) {
	p.X *= s
	p.Y *= s
}
func (p *Point) Abs() float64 {
	return math.Sqrt(float64(p.X*p.X + p.Y*p.Y))
}
func main() {
	p := new(Point)
	p.X = 3
	p.Y = 4
	m = p
	fmt.Printf("The length of the vector p is: %f\n", m.Abs())
	fmt.Printf("Point p has the coordinates: X %f - Y %f\n", p.X, p.Y)
	p.Scale(5)
	fmt.Printf("After scaling:\n%f\n", m.Abs())
	fmt.Printf("X %f - Y %f\n", p.X, p.Y)
}
