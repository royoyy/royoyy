package main

import "fmt"

func main() {
	s := []int{1, 2, 3}
	fmt.Println(s)
	fmt.Println("The length of s before enlarging is:", len(s))
	factor := 5
	s = enlarge(s, factor)
	fmt.Println(s)
	fmt.Println("The length of s after enlarging is:", len(s))
}

func enlarge(s []int, factor int) []int {
	ns := make([]int, len(s)*factor)
	// fmt.Println("The length of ns  is:", len(ns))
	copy(ns, s)
	//fmt.Println(ns)
	s = ns
	//fmt.Println(s)
	//fmt.Println("The length of s after enlarging is:", len(s))
	return s
}
