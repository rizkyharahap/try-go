package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	// address1 := Address{"Subang", "Jawa Barat", "Indonesia"}
	// address2 := address1 // copy value not reference

	// address2.City = "Bandung"
	// fmt.Println(address1) // not change
	// fmt.Println(address2) // change to Bandung

	var address1 Address = Address{"Subang", "Jawa Barat", "Indonesia"}
	var address2 *Address = &address1 // pointer

	address2.City = "Bandung"
	fmt.Println(address1) // changed
	fmt.Println(address2) // changed
}
