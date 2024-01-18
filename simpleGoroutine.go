package main


import (
	"fmt"
	"runtime"
)

func print(til int, message string) {
	for i := 0; i < til; i++ {
		fmt.Println((i+1), message)
	}
}

func main() {
	runtime.GOMAXPROCS(2)

	go print(10, "halo")
	print(10, "apa kabar")

	var input string
	fmt.Scanln(&input)
	
}