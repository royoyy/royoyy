package main

import "fmt"

type Rectangle struct {
	w, h int
}

func Area(r *Rectangle) int {
	return r.w * r.h
}
func Perimeter(r *Rectangle) int {
	return 2 * (r.w + r.h)
}
func main() {
	r := &Rectangle{3, 4}
	fmt.Println(r)
	fmt.Println("area:", Area(r))
	fmt.Println("perimeter:", Perimeter(r))
}
