package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	// incr := func(wg *sync.WaitGroup) {s
	// 	var i int
	// 	wg.Add(1)
	// 		defer wg.Done()
	// 		i++
	// 		fmt.Printf("value of i : %v\n", i)
	// }

	for i := 1; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fmt.Println(i)
		}(i)
	}

	wg.Wait()
}
