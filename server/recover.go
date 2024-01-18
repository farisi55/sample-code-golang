package main

import "fmt"

func recoverFullName() {
	if r := recover(); r != nil {
		fmt.Println("recover from ", r)
	} 
}

