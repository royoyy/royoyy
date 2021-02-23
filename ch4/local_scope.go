package main

import "fmt"

var a = "G" // global (package) scope
func main() {
	n()
	m()
	n()
	fmt.Println()
}
func n() {
	fmt.Print(a)
}
func m() {
	a := "O"
	fmt.Print(a)
}
