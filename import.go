package main

import (
	"fmt"
	"kelas-golang-pzn/go-dasar/helper"
)

func main() {

	result := helper.SayHello("rafli")
	fmt.Println(result)

	fmt.Println(helper.Application)
	fmt.Println(helper.version)
	fmt.Println(helper.sayGoodbye("rafli"))
}
