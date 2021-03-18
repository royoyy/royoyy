package main

import "fmt"

type Simpler interface {
	Get() int
	Put(int)
}
type Simple struct {
	i int
}

func (p *Simple) Get() int {
	return p.i
}
func (p *Simple) Put(u int) {
	p.i = u
}

type Any interface{}

func gI(any interface{}) int {
	if v, ok := any.(Simpler); ok {
		return v.Get()
	}
	return 0
}
func main() {
	s := &Simple{5}
	fmt.Println(gI(s), "setted")
}
