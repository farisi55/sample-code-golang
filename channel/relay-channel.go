package main

import (
	"fmt"
)

func genMsg(ch1 chan<- string) {
	//send message on ch1
	ch1 <- "message"
}

func relayMsg(ch1 <-chan string, ch2 chan<- string) {
	//revc message on ch1
	m := <-ch1
	//send it on ch2
	ch2 <- m
}

func main() {
	//create ch 1 and ch 2
	ch1 := make(chan string)
	ch2 := make(chan string)

	//spine gorothine genMsg and relayMsg
	go genMsg(ch1)
	go relayMsg(ch1, ch2)

	//recv message on ch2
	v := <-ch1
	fmt.Println(v)

}
