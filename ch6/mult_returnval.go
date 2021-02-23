package main

import "fmt"

func main() {
	n1, n2 := 3, 4
	fmt.Println("ints:", n1, n2)
	sum, prod, diff := sumProductDiff(n1, n2)
	fmt.Println("named  : Sum:", sum, "| Product:", prod, "| Diff:", diff)
	sum, prod, diff = sumProductDiffN(n1, n2)
	fmt.Println("unamed : Sum:", sum, "| Product:", prod, "| Diff:", diff)
}
func sumProductDiff(a, b int) (int, int, int) {
	return a + b, a * b, a - b
}
func sumProductDiffN(a, b int) (s, p, d int) {
	s, p, d = a+b, a*b, a-b
	return
}
