package main

import "fmt"

func main() {
	name := "muhammad"

	if name == "muhammad" {
		fmt.Println("halo, muhammad")
	} else if name == "rafli" {
		fmt.Println("halo, rafli")
	} else {
		fmt.Println("hai, boleh kenalan")
	}

	if length := len(name); length > 5 {
		fmt.Println("nama terlalu panjang")
	} else {
		fmt.Println("nama sudah benar")
	}
}
