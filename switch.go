package main

import "fmt"

func main() {
	name := "rafli"

	switch name {
	case "rafli":
		fmt.Println("welcome")
	case "muhammad":
		fmt.Println("assalamu`alaikum")
	default:
		fmt.Println("hi, boleh kenalan?")
	}

	switch length := len(name); length > 5 {
	case true:
		fmt.Println("nama terlalu panjang")
	case false:
		fmt.Println("nama sudah benar")
	}

	name = "muhammad"

	length := len(name)
	switch {
	case length > 10:
		fmt.Println("nama sudah benar")
	case length > 5:
		fmt.Println("nama terlalu panjang")
	default:
		fmt.Printf("nama sudah benar: %s", name)
	}
}
