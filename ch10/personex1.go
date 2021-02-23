package main

import (
	"fmt"
	"strings"
)

type Person struct {
	firstname string
	lastname  string
}

func upPerson(p Person) {
	p.firstname = strings.ToUpper(p.firstname)
	p.lastname = strings.ToUpper(p.lastname)
}
func main() {
	var pers1 Person
	pers1.firstname = "Chris"
	pers1.lastname = "Woodward"
	upPerson(pers1)
	fmt.Printf("The name of the person is %s %s\n", pers1.firstname, pers1.lastname)
}
