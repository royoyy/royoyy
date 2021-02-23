package main

import "fmt"

func main() {
	l := 5
	fibs := fibonacciS(l)
	fmt.Println("the frist", l, "fibonacci numbers:\n", fibs)
}
func fibonacciS(l int) []int {
	fibs := make([]int, l)
	fibs[0], fibs[1] = 1, 1
	for i := 2; i < l; i++ {
		fibs[i] = fibs[i-1] + fibs[i-2]
	}
	return fibs
}
