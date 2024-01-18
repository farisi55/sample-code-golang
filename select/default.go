package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)

	go func() {
		for i := 0; i < 3; i++ {
			time.Sleep(1 * time.Second)
			ch <- "message"
		}
	}()

	//TODO if thre is no value on channel , do not block

	for i := 0; i < 2; i++ {
		select {
		case msg := <-ch:
			fmt.Println(msg)
		default:
			fmt.Println("no message")
		}
	}

	//Do some processing
	fmt.Println("processing ... ")
	time.Sleep(2000 * time.Millisecond)
}
