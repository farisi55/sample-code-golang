package main

import (
	"fmt"
	"runtime"
)

func main() {

	runtime.GOMAXPROCS(4)

	ch := make(chan int, 3)

	go func() {
		defer close(ch)
		for i := 0; i < 20; i++ {
			fmt.Println("sending : ", i)
			ch <- i
		}
	}()

	for v := range ch {
		fmt.Println("Receiving : ", v)
	}
}
