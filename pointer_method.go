package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Merried() {
	man.Name = "Mr. " + man.Name
}

func main() {

	john := Man{"john wick"}
	john.Merried()
	fmt.Print(john.Name)
}
