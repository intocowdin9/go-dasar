package main

import "fmt"

func main() {
	// var person map[string]string = map[string]string{}
	// person["name"] = "muhammad"
	// person["province"] = "aceh"

	person := map[string]string{
		"name":     "muhammad",
		"province": "aceh",
	}

	fmt.Println(person["name"])
	fmt.Println(person["province"])
	fmt.Println(person)

	book := make(map[string]string)
	book["title"] = "belajar go dasar bersama programmer zaman now"
	book["author"] = "muhammad rafli"
	book["wrong"] = "ups"

	delete(book, "wrong")

	fmt.Println(book)

}
