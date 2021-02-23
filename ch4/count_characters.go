package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	str1 := "hello, world"
	fmt.Println("original string:", str1)
	fmt.Printf("the number of bytes in string str1 is %d\n", len(str1))
	fmt.Printf("the number of characters in string str1 is %d\n", utf8.RuneCountInString(str1))
	str2 := "hello, world汉字"
	fmt.Println("original string:", str2)
	fmt.Printf("...%d\n", len(str2))
	fmt.Printf("...%d\n", utf8.RuneCountInString(str2))
}
