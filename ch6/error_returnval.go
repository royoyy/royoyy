package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	fmt.Println("First example with -1:")
	ret, err := mySqrt(-1)
	if err != nil {
		fmt.Println("Error! Return values are: ", ret, err)
	} else {
		fmt.Println("It's OK! Return values are: ", ret, err)
	}
	fmt.Println("Second example with 5:")
	ret, err = mySqrt(5)
	if err != nil {
		fmt.Println("...", ret, err)
	} else {
		fmt.Println("...", ret, err)
	}
	ret, err = mySqrtN(5)
	fmt.Println("...", ret, err)
}
func mySqrt(f float64) (float64, error) {
	if f < 0 {
		return float64(math.NaN()), errors.New("Got a negative number")
	}
	return math.Sqrt(f), nil
}
func mySqrtN(f float64) (ret float64, err error) {
	if f < 0 {
		ret = float64(math.NaN())
		err = errors.New("Got a negative number")
	} else {
		ret = math.Sqrt(f)
		err = nil
	}
	return
}
