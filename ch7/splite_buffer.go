package main

import (
	"bytes"
	"fmt"
)

func main() {
	var buf bytes.Buffer
	buf.WriteString("buffer")
	str := buf.String()
	fmt.Println("the string before splitting is:", str)
	str1, str2 := str[:3], str[3:]
	fmt.Println("the strings after splitting are:", str1, str2)
}
