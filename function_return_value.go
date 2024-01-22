package main

import "fmt"

func getHello(name string) string {
	return "halo, " + name
}

func main() {
	result := getHello("rafli")
	fmt.Println(result)
}
