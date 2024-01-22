package main

import "fmt"

func getCompleteName() (firstName, middleName, lastName string) {
	firstName = "muhammad"
	middleName = "rafli"
	lastName = "gaspul"

	return firstName, middleName, lastName
}

func main() {
	a, b, c := getCompleteName()
	fmt.Println(a, b, c)
}
