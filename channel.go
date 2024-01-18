package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(2)

	var messages = make(chan string)
	
	var sayHelloTo = func(who string) {
		var data = fmt.Sprintf("hello %s", who)
		messages <- data
	}

	go sayHelloTo("john wick")
	go sayHelloTo("ethan hunt")
	go sayHelloTo("jason bourne")

	var messages1 = <-messages
	fmt.Println(messages1)

	var messages2 = <-messages
	fmt.Println(messages2)

	var messages3 = <-messages
	fmt.Println(messages3)
}