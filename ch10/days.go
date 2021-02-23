package main

import "fmt"

type Day int

var dayName = []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}

const (
	MO Day = iota
	TU
	WE
	TH
	FR
	SA
	SU
)

func (d Day) String() string {
	return dayName[d]
}
func main() {
	var d Day = 3
	fmt.Printf("The 3rd day is: %s\n", d)
	var day = SU
	fmt.Println(day)
	fmt.Println(0, MO, 1, TU)
}
