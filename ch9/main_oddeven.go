package main

import (
	"./even"
	"fmt"
)

func main() {
	for i := 0; i <= 10; i++ {
		fmt.Printf("Is the integer %d even?\n%v\n", i, even.IsEven(i))
	}
}
