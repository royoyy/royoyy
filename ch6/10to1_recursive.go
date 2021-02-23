package main

import "fmt"

var tar int = 10

func main() {
	pos := 1
	printrec(pos)
}
func printrec(pos int) {
	if pos > tar {
		return
	}
	printrec(pos + 1)
	fmt.Printf("%d\n", pos)
}
