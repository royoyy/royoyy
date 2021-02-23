package main

import "fmt"

func main() {
	pos := 4
	res := fibonacci2(pos)
	fmt.Printf("the %d-th(from 0-th) fibonacci number is %d\n", pos, res)
	pos = 10
	res = fibonacci2(pos)
	fmt.Printf("the %d-th(from 0-th) fibonacci number is %d\n", pos, res)
}
func fibonacci2(pos int) (res int) {
	if pos < 2 {
		res = 1
	} else {
		a := fibonacci2(pos - 1)
		b := fibonacci2(pos - 2)
		res = a + b
	}
	return
}
