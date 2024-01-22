package main

import "fmt"

func main() {
	// const firstName string = "muhammad"
	// const lastName = "rafli"

	const (
		firstName string = "muhammad"
		lastName         = "rafli"
	)

	lastName = "halo" //error
	fmt.Println(firstName)
	fmt.Println(lastName)
}
