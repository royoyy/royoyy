package main

import "fmt"

func main() {
	arr := [3]float32{1, 2, 3}
	fmt.Println("sum of", arr, "is:", sumArray(arr))
	sl1 := []float32{1, 2, 3}
	fmt.Println("sum of", sl1, "is:", sumSlice(sl1))
	sl2 := []int{1, 2, 3}
	sum, ave := sumAndAverage(sl2)
	fmt.Println("sum of", sl2, "is:", sum)
	fmt.Println("average of", sl2, "is:", ave)
}
func sumArray(arr [3]float32) float32 {
	res := float32(0)
	for _, f := range arr {
		res += f
	}
	return res
}
func sumSlice(sl []float32) float32 {
	res := float32(0)
	for _, f := range sl {
		res += f
	}
	return res
}
func sumAndAverage(sl []int) (int, float32) {
	sum := 0
	for _, n := range sl {
		sum += n
	}
	return sum, float32(sum / len(sl))
}
