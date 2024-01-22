package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {

	var address1 Address = Address{"subang", "jawa barat", "indonesia"}
	var address2 *Address = &address1 // pointer

	address2.City = "bandung"
	fmt.Println(address1) // ikut berubah
	fmt.Println(address2) // berubah

	// address2 = &Address{"panton labu", "aceh timur", "indonesia"}
	*address2 = Address{"panton labu", "aceh timur", "indonesia"}
	fmt.Println(address1)
	fmt.Println(address2)
}
