package main

import "fmt"

type S struct {
	f float32
	int
	string
}

func main() {
	s := &S{1.5, 1, "a string"}
	fmt.Println(s.int)
	fmt.Println(s.string)
}
