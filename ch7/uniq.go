package main

import "fmt"

func main() {
	arr := [5]byte{'a', 'a', 'a', 'd', 'f'}
	fmt.Println(arr)
	uni := make([]byte, len(arr))
	uni[0] = arr[0]
	count, t := 1, arr[0]
	for _, b := range arr {
		if b != t {
			uni[count] = b
			count++
		}
		t = b
	}
	uni = uni[:count]
	fmt.Println(uni)
}
