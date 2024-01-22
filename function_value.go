package main

import "fmt"

func getGoodbye(name string) string {
	return "good bye, " + name
}

func main() {
	goodbye := getGoodbye
	fmt.Println(goodbye("rafli"))
}
