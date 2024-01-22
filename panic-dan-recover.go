package main

import "fmt"

func endApp() {
	fmt.Println("end app")
	message := recover()
	fmt.Println("terjadi panic", message)
}

func runApp(error bool) {
	defer endApp()

	if error {
		panic("ups error")
	}
}

func main() {
	runApp(true)
	fmt.Println("bye..")
}
