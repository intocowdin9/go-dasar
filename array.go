package main

import "fmt"

func main() {
	var names [2]string
	names[0] = "muhammad"
	names[1] = "rafli"

	fmt.Println(names[0])
	fmt.Println(names[1])
	// fmt.Println(names[3]) //error out of bounds

	var values = [3]int{
		90,
		80,
		95,
	}

	fmt.Println(values)

	var valuex = [...]int{
		90,
		80,
		95,
		95,
		95,
	}

	fmt.Println(len(valuex))
}
