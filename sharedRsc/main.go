package main

import (
	"fmt"
	"sync"
	"time"
)

var sharedRsc = make(map[string]interface{})

func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()

		//TODO: suspend goroutine until sharedRsc is populated.

		for len(sharedRsc) == 0 {
			time.Sleep(1 * time.Millisecond)
		}

		fmt.Println("Result: ", sharedRsc["rsc1"])
	}()

	// writes changes to sharedRsc
	fmt.Println("Data assigment for on Other Go Routine")
	sharedRsc["rsc1"] = "Data value assign by main Gorotine"

	wg.Wait()
}
