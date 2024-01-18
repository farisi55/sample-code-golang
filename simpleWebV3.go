package main

import (
	"fmt"
	"net/http"
)

func main() {
	handlerIndex := func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello nama orang"))
	}

	http.HandleFunc("/", handlerIndex)
	http.HandleFunc("/index", handlerIndex)
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello Lagi dan lagi nama orang"))
	})

	fmt.Printf("Server started at port :9000")

	http.ListenAndServe(":9000", nil)	
}