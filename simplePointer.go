package main

import "fmt"

func main(){
	var numberA int = 4
	var numberB *int = &numberA

	fmt.Println("numberA (value): ", numberA) //4
	fmt.Println("numberA (address): ", &numberA) // output address nya 0xc...
	fmt.Println("numberB (value): ", *numberB) //4
	fmt.Println("numberB (address): ", numberB) // output address nya 0xc...


	fmt.Println("before change :", numberA)
	change(&numberA, 10)
	fmt.Println("after change :", numberA)
}

func change(original *int, value int) {
	*original = value
}