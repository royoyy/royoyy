package main

import (
	"./stack"
	"fmt"
)

func main() {
	s := new(stack.Stack)
	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Push(1)
	s.Push(2)
	s.Push(3)
	fmt.Println(s)
	p := s.Pop()
	fmt.Println(p, "popped")
	p = s.Pop()
	fmt.Println(p, "popped")
	p = s.Pop()
	fmt.Println(p, "popped")
	p = s.Pop()
	fmt.Println(p, "popped")
	p = s.Pop()
	fmt.Println(p, "popped")
	fmt.Println(s)
}
