package main

import "fmt"

type Address struct {
	country string
	state   string
	city    string
	detail  string
}
type VCard struct {
	name string
	addr *Address
	bir  string
	pro  string
}

func main() {
	ad := &Address{"us", "ca", "la", "detail"}
	vc := VCard{"trump", ad, "2016", "president"}
	fmt.Println(vc)
	fmt.Println("address:", vc.addr)
}
