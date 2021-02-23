package main

import "fmt"

func main() {
	w, h := 10, 5
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
