package main

import "fmt"

func main() {
	// use 2 nested for loops
	for i := 1; i <= 5; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("G")
		}
		fmt.Println()
	}
	// use only one for loop and string concatenation
	str := "G"
	for i := 1; i <= 5; i++ {
		fmt.Println(str)
		str += "G"
	}
}
