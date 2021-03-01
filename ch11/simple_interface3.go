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
	// return any.(Simpler).Get()   // unsafe, runtime panic possible
	if v, ok := any.(Simpler); ok {
		return v.Get()
	}
	return 0 // default value
}
func main() {
	s := &Simple{1}
	fmt.Println(gI(s))
}
