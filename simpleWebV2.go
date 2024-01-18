package main

import (
	"fmt"
	"net/http"
	"time"
)

func handlerIndex(w http.ResponseWriter, r *http.Request) {
	var message = "Welcome nama orang"
	w.Write([]byte(message))
}

func handlerHello(w http.ResponseWriter, r *http.Request) {
	var message = "Hello XL dan nama orang"
	time.Sleep(11000 * time.Millisecond)
	w.Write([]byte(message))
}

func main() {
	http.HandleFunc("/", handlerIndex)
	http.HandleFunc("/index", handlerIndex)
	http.HandleFunc("/hello", handlerHello)

	var address = "localhost:9000"
	fmt.Printf("Server started at %s\n", address)

	server := new(http.Server)
	server.Addr = address
	server.ReadTimeout = 10 * time.Second
	server.WriteTimeout = 10 * time.Second

	err := http.ListenAndServe(address, nil)
	if err != nil {
		fmt.Println(err.Error())
	}

	
}