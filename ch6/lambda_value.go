package main

import "fmt"

func main() {
	fv := func() {
		fmt.Println("Hello World!")
	}
	fv()
	fmt.Printf("the type of fv is %T\n", fv)
}
