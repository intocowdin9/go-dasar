package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func main() {
	var rafli Customer

	rafli.Name = "muhammad rafli"
	rafli.Address = "aceh"
	rafli.Age = 26

	fmt.Println(rafli)
	fmt.Println(rafli.Name)
	fmt.Println(rafli.Address)
	fmt.Println(rafli.Age)

	// struct literal
	acehWoman := Customer{
		Name:    "fatimah",
		Address: "aceh utara",
		Age:     21,
	}

	fmt.Println(acehWoman)

	apaAceh := Customer{Name: "bregeh", Address: "kruengkreh", Age: 43}

	fmt.Println(apaAceh)
}
