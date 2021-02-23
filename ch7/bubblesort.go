package main

import "fmt"

func main() {
	sl := []int{12, 13, 14, 11, 1, 2, 3}
	for i := 0; i < len(sl); i++ {
		for j := 0; j < len(sl)-i-1; j++ {
			if sl[j] > sl[j+1] {
				sl[j], sl[j+1] = sl[j+1], sl[j]
			}
		}
	}
	fmt.Println(sl)
}
