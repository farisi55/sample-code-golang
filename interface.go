package main

import (
	"fmt"
)

type HasName interface {
	GetName() string
}

func SayHello(hashName HasName) {
	fmt.Println("Hello", hashName.GetName())
}

type Person struct {
	Name string 
}

func (person Person) GetName() string {
	return person.Name
}

type Animal struct {
	Name string
}

func (animal Animal) GetName() string {
	return animal.Name
}

func main() {
	var feri Person
	feri.Name = "feri"

	SayHello(feri)

	cat := Animal{
		Name: "push",
	}

	SayHello(cat)
}