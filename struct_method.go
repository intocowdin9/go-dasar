package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func (customer Customer) sayHello(name string) {
	fmt.Println("halo", name, "my name is", customer.Name)
}

func main() {
	wahyu := Customer{Name: "wahyu umaya", Address: "aceh", Age: 25}
	wahyu.sayHello("rafli")

}
