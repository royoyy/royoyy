package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}
type Polar struct {
	l, r float64
}

func (p *Point) Abs() float64 {
	return math.Sqrt(p.x*p.x + p.y*p.y)
}
func (p *Polar) Scale(f float64) {
	p.l *= f
}
func main() {
	p1 := &Point{1.0, 2.0}
	fmt.Println(p1.Abs())
	p2 := &Polar{3.0, 4.0}
	p2.Scale(5.0)
	fmt.Println(p2.l)
}
