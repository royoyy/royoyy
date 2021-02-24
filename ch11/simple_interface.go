package main

import "fmt"

type Simpler interface {
	Get() int
	Set(n int)
}
type Simple struct {
	n int
}

func (s Simple) Get() int {
	return s.n
}
func (s *Simple) Set(n int) {
	s.n = n
}
func testSimpler(s Simpler) {
	s.Set(2)
	fmt.Println("Get():", s.Get())
}
func main() {
	s := &Simple{1}
	fmt.Println(s)
	testSimpler(s)
}
