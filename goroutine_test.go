package main

import (
	"fmt"
	"testing"
	"time"
)

func RunHelloWorld() {
	fmt.Println("Hello world")
}

func TestCreateGoroutine(t *testing.T) {
	go RunHelloWorld()
	fmt.Println("hay")

	time.Sleep(1 * time.Second)
}

func DisplayNumber(number int) {
	fmt.Println("Diplay: ", number)
}

func TestManyGoroutines(t *testing.T) {
	for i := 0; i < 100; i++ {
		go DisplayNumber(i)
	}
}
