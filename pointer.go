package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {

	// address1 := Address{"lhokseumawe", "aceh", "indonesia"}
	// address2 := address1 //copy value

	// address2.City = "bandung"
	// fmt.Println(address1) // tidak berubah
	// fmt.Println(address2) // berubah

	var address1 Address = Address{"lhokseumawe", "aceh", "indonesia"}
	var address2 *Address = &address1 // pointer

	address2.City = "bandung"
	fmt.Println(address1) // ikut berubah
	fmt.Println(address2) // berubah
}
