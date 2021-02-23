package main

import "fmt"

func main() {
	sl := []int{4, 9, 7}
	fmt.Println("Min:", minSlice(sl), "Max:", maxSlice(sl))
}
func minSlice(sl []int) int {
	min := sl[0]
	for _, n := range sl {
		if n < min {
			min = n
		}
	}
	return min
}
func maxSlice(sl []int) int {
	max := sl[0]
	for _, n := range sl {
		if n > max {
			max = n
		}
	}
	return max
}
