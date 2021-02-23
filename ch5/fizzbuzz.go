package main

import "fmt"

const (
	FIZZ     = 3
	BUZZ     = 5
	FIZZBUZZ = 15
)

func main() {
	for i := 0; i <= 15; i++ {
		switch {
		case i%FIZZ == 0:
			fmt.Println("FIZZ")
		case i%BUZZ == 0:
			fmt.Println("BUZZ")
		case i%FIZZBUZZ == 0:
			fmt.Println("FIZZBUZZ")
		default:
			fmt.Println(i)
		}
	}
}
