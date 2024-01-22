package main

import "fmt"

func Ups() any {
	// return 1
	// return true
	return "ups"
}

func main() {

	var kosong any = Ups()
	fmt.Print(kosong)
}
