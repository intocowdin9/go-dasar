package main

import "fmt"

func random() any {
	return "OK"
}

func main() {

	var result any = random()

	// var resultString string = result.(string)
	// fmt.Println(resultString)

	// var resultInt int = result.(int) // error conversion
	// fmt.Println(resultInt)

	// using switch

	switch value := result.(type) {
	case string:
		fmt.Println("string", value)
	case int:
		fmt.Println("int", value)
	default:
		fmt.Print("unknow")
	}
}
