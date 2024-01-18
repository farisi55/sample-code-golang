package main

import "fmt"

func fullName(firstName *string, lastName *string) {
	defer fmt.Println("deferred call  in fullname")

	if firstName == nil {
		panic("runtime error: first name cannot be nill")
	}

	if lastName == nil {
		panic("runtime error: last name cannot be nill")
	}

	fmt.Printf("%s %s\n", *firstName, *lastName)
	fmt.Println("returned normally from fullName")
}

func main() {
	defer fmt.Println("deferred call in main")
	firstName := "Elon"
	lastName := "Mask"
	fullName(&firstName, &lastName)
	fmt.Println("returned normally from main")
}
