package main

import "fmt"

func main() {
	sl := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	fun := func(n int) int { return n * 10 }
	fmt.Println(mapFunc(fun, sl))
}
func mapFunc(fun func(int) int, sl []int) []int {
	nsl := make([]int, len(sl))
	for i, n := range sl {
		nsl[i] = fun(n)
	}
	return nsl
}
