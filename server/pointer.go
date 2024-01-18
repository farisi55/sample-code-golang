package main

import "fmt"

func tesPointer(valA *string, valB *int) {
	fmt.Println("alamat string argument ", valA)
	fmt.Println("alamat Int argument ", valB)

	fmt.Println("value first argument ", *valA)
	fmt.Println("value second argument ", *valB)

	*valA = "first argument change"
	*valB = 12
}

func testValue(valA string, valB int) {
	valA = "first arg changed in testValue Function"
	valB = 11
}

func main() {
	s := "This is string argument"
	i := 10

	tesPointer(&s, &i)
	testValue(s, i)

	fmt.Println("test Value ==> ", testValue)
	fmt.Println("test Pointer ==>", tesPointer)
}
