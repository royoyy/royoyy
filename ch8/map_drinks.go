package main

import (
	"fmt"
	"sort"
)

func main() {
	drinks := map[string]string{"coke": "可乐", "juice": "果汁", "tea": "茶"}
	fmt.Println(drinks)
	sl := make([]string, len(drinks))
	i := 0
	for key := range drinks {
		sl[i] = key
		i++
	}
	sort.Strings(sl)
	for _, key := range sl {
		fmt.Println(key, drinks[key])
	}
}
