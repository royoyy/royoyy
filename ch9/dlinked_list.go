package main

import (
	"container/list"
	"fmt"
)

func main() {
	l := list.New()
	l.PushBack(101)
	l.PushBack(102)
	l.PushBack(103)
	fmt.Println(l)
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e, e.Value)
	}
}
