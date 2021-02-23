package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y float64
}
type Polar struct {
	l, r float64
}

func Abs(p *Point) float64 {
	return math.Sqrt(p.x*p.x + p.y*p.y)
}
func Scale(p *Polar, f float64) {
	p.l *= f
}
func main() {
	p := &Point{1, 2}
	fmt.Println(Abs(p))
	q := &Polar{3, 4}
	Scale(q, 5)
	fmt.Println(q.l)
}
