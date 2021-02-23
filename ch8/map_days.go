package main

import "fmt"

func main() {
	days := map[int]string{0: "Sunday", 1: "Monday", 2: "Tuesday", 3: "Wednesday", 4: "Thursday", 5: "Friday", 6: "Saturday"}
	for key := range days {
		fmt.Println(key, days[key])
	}
}
