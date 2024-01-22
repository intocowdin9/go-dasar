package main

import "fmt"

func sumAll(numbers ...int) int {
	total := 0

	for _, number := range numbers {
		total += number
	}

	return total
}

func main() {
	fmt.Println(sumAll(3, 3, 3))
	fmt.Println(sumAll(3, 3, 3, 3, 3))

	numbers := []int{3, 3, 3}
	fmt.Println(sumAll(numbers...))
}
