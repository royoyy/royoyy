package main

import "fmt"

func main() {
	str := "asdf汉字"
	fmt.Println(str, reverseString(str))
}
func reverseString(s string) string {
	/* ascii only
	str := []byte(s)
	for i, j := 0, len(str)-1; i < len(str)/2; i, j = i+1, j-1 {
		t := str[i]
		str[i] = str[j]
		str[j] = t
	}
	return string(str)
	*/
	runes := []rune(s)
	for i := 0; i < len(runes)/2; i++ {
		runes[i], runes[len(runes)-i-1] = runes[len(runes)-i-1], runes[i]
	}
	return string(runes)
}
