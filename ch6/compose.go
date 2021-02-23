package main

import (
	"fmt"
	"math"
)

func main() {
	com := compose(math.Sin, math.Cos)
	fmt.Println(com(0.5))
	// another way
	fmt.Println(compose(math.Sin, math.Cos)(0.5))
}
func compose(f, g func(x float64) float64) func(x float64) float64 {
	return func(x float64) float64 {
		return f(g(x))
	}
}
