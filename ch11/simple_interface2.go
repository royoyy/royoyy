package main

import "fmt"

type Simpler interface {
	Get() int
	Put(int)
}
type Simple struct {
	i int
}
type RSimple struct {
	ri int
}

func (p *Simple) Get() int {
	return p.i
}
func (p *Simple) Put(u int) {
	p.i = u
}

func (rp *RSimple) Get() int {
	return rp.ri
}
func (rp *RSimple) Put(ru int) {
	rp.ri = ru
}
func fI(s Simpler) {
	switch s.(type) {
	case *Simple:
		fmt.Println("Simple")
	case *RSimple:
		fmt.Println("RSimple")
	case nil:
		fmt.Println("nil")
	default:
		fmt.Println("unexpected")
	}
}
func main() {
	s := &Simple{1}
	fI(s)
}
