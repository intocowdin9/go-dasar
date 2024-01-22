package main

import "fmt"

func sayHelloTo(firstName string, lastName string) {
	fmt.Println("halo,", firstName, lastName)
}

func main() {
	sayHelloTo("muhammad", "rafli")
}
