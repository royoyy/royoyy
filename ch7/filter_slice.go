package main

import "fmt"

func main() {
	sl := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(filter(sl, isEven))
}
func filter(sl []int, fn func(n int) bool) []int {
	nsl := make([]int, len(sl))
	at := 0
	for _, n := range sl {
		if fn(n) {
			nsl[at] = n
			at++
		}
	}
	return nsl[0:at]
}
func isEven(n int) bool {
	if n%2 == 0 {
		return true
	}
	return false
}
