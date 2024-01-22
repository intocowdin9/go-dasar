package main

import "fmt"

type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("you are blocked!", name)
	} else {
		fmt.Println("welcome", name)
	}
}

func main() {

	blacklist := func(name string) bool {
		return name == "anjing"
	}

	registerUser("rafli", blacklist)

	registerUser("anjing", func(name string) bool {
		return name == "anjing"
	})
}
