package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		ch <- "one"
	}()

	//TODO timeout

	select {
	case msg := <-ch:
		fmt.Println(msg)
	case <-time.After(time.Second * 1):
		fmt.Println("timeout")
	}
}
