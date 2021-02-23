package main

import "fmt"

type employee struct {
	salary float32
}

func (recv employee) giveRaise(rate float32) float32 {
	return recv.salary * (1 + rate)
}
func main() {
	e := &employee{10}
	fmt.Println("original salary:", e.salary)
	r := 0.5
	fmt.Println("salary raised to", e.giveRaise(float32(r)))
}
