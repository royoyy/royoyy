package main

import "fmt"

var a string // global scope

func main() {
	a = "G"
	fmt.Print(a)
	f1()
	fmt.Println()
}
func f1() {
	a := "O"
	fmt.Print(a)
	f2()
}
func f2() {
	fmt.Print(a)
}
