package main

import "fmt"

func main() {
	sl := []string{"a", "b", "c", "d", "e"}
	fmt.Println("the slice before removing is:", sl)
	start, end := 2, 4
	fmt.Println("the slice after removing is:", removeStringSlice(sl, start, end))
}
func removeStringSlice(sl []string, start int, end int) []string {
	for i, j := start, end; i < end && j < len(sl); i, j = i+1, j+1 {
		sl[i] = sl[j]
	}
	nlen := len(sl) - (end - start)
	nsl := sl[:nlen]
	return nsl
}
