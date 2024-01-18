package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var data int

	wg.Add(1)
	go func() {
		defer wg.Done()
		data++
	}()

	//wg.Add(1)
	go func() {
		//defer wg.Done()
		data++
	}()

	wg.Wait()
	fmt.Printf("the value of data %v\n", data)

}
