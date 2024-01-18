package main

import "fmt"

type Customer struct {
	Name, Address string
	Age 			int
}

func (customer Customer) sayHi(name string) {
	fmt.Println("Hello", name, "My Name is", customer.Name)
}

func (a Customer) sayHuu() {
	fmt.Println("Huu from", a.Name)
}

func main() {
	var feri Customer
	feri.Name = "feri"
	feri.Address = "Indonesia"
	feri.Age = 30

	feri.sayHi("Joko")
	feri.sayHuu()
}