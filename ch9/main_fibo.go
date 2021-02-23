package main

import (
	"./fibo"
	"fmt"
)

func main() {
	n := 5
	fmt.Printf("the %d-th fabonacci number is %d\n", n, fibo.Fibonacci(n))
}
