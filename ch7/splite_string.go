package main

import "fmt"

func main() {
	str := "abcd"
	fmt.Println("the string before splitting is:", str)
	pos := 2
	str1, str2 := spliteString(str, pos)
	fmt.Println("the strings after splitting are:", str1, str2)
}
func spliteString(str string, pos int) (string, string) {
	return str[:pos], str[pos:]
}
