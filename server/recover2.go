package main

import "fmt"

func handlePanic() {
	if a := recover(); a != nil {
		fmt.Println("Recover: ", a)
	}
}

func entry(lang *string, aname *string) {
	defer handlePanic()

	if lang == nil {
		panic("Error: languge cannot be nil")
	}

	if aname == nil {
		panic("Error: Author cannot be nil")
	}

	fmt.Printf("Author languge: %s\n  Author name: %s\n", lang, aname)
	fmt.Println("Return successfuly from the entry function")
}

func main() {
	A_lang := "Go Langunge"
	entry(&A_lang, nil)
	fmt.Println("Return successfuly from the main function")
}
