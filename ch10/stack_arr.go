package main

import "fmt"

const LIMIT = 4

type Stack [LIMIT]int

func (s *Stack) Push(n int) {
	for i, v := range s {
		if v == 0 {
			s[i] = n
			break
		}
	}
}
func (s *Stack) Pop() (n int) {
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] != 0 {
			n = s[i]
			s[i] = 0
			break
		}
	}
	return
}
func (s Stack) String() string {
	str := ""
	for i, v := range s {
		str += fmt.Sprintf("[%d:%d] ", i, v)
	}
	return str
}
func main() {
	s := new(Stack)
	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Push(4)
	s.Push(5)
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
