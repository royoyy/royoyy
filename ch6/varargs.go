package main

import "fmt"

func main() {
	printInts(2, 3)
	fmt.Println()
	printInts(0, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println()
}
func printInts(list ...int) {
	for _, v := range list {
		fmt.Printf("%d ", v)
	}
}
