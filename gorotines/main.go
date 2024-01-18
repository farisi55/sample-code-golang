package main

import (
	"fmt"
	"time"
)

func fun(s string) {
	for i := 0; i < 3; i++ {
		fmt.Println(s)
	}
}

func main() {
	fun("direct all")

	go fun("gorouthine-1")

	go func() {
		fun("gorouthine-2")
	}()

	fv := fun
	go fv("gorouthine-3")

	time.Sleep(time.Second * 1)

	fmt.Println("wait for gorothine ...")
	fmt.Println("done ...")
}
