package main

import "fmt"

func main() {
	type NoKTP string

	var ktpRafli NoKTP = "7777777"

	var contoh string = "333333333"
	var contohKtp NoKTP = NoKTP(contoh)

	fmt.Println(ktpRafli)
	fmt.Println(contohKtp)
}
