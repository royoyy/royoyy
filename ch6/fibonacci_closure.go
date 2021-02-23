package main

import "fmt"

func main() {
	fib := fibonacciC()
	fmt.Println(fib(), fib(), fib())
}
func fibonacciC() func() int {
	a, b := 1, 1
	return func() int {
		a, b = b, a+b
		return b
	}
}
