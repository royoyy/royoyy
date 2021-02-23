package main

import "fmt"

type flt func(int) bool

func isEven(n int) bool {
	if n%2 == 0 {
		return true
	}
	return false
}

func main() {
	slice := []int{1, 2, 3, 4, 5, 7}
	fmt.Println("slice:", slice)
	even, odd := filter(slice, isEven)
	fmt.Println("the even elements of slice are:", even)
	fmt.Println("the odd elements of slice are:", odd)
}
func filter(sl []int, f flt) (is, not []int) {
	for _, val := range sl {
		if f(val) {
			is = append(is, val)
		} else {
			not = append(not, val)
		}
	}
	return
}
