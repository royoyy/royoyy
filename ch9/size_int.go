package main

import (
	"fmt"
	"unsafe"
)

func main() {
	fmt.Println("the size of an int is:", unsafe.Sizeof(int(0)))
}
