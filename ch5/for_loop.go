package main

import "fmt"

func main() {
	// use for
	for i := 0; i < 5; i++ {
		fmt.Printf("the counter is at %d\n", i)
	}
	// use goto
	i := 0
START:
	fmt.Printf("the counter is at %d\n", i)
	i++
	if i < 5 {
		goto START
	}
}
