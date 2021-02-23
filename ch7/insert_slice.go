package main

import "fmt"

func main() {
	sl1 := []string{"a", "b", "e"}
	sl2 := []string{"c", "d"}
	pos := 2
	fmt.Println(insertStringSlice(sl1, sl2, pos))
}
func insertStringSlice(sl1 []string, sl2 []string, pos int) []string {
	nsl := make([]string, len(sl1)+len(sl2))
	copy(nsl, sl1[:pos])
	at := copy(nsl[pos:], sl2)
	copy(nsl[pos+at:], sl1[pos:])
	return nsl
}
