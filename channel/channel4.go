package main

import "fmt"

func main() {

	owner := func() chan int {
		ch := make(chan int, 5)

		go func() {
			defer close(ch)
			for i := 0; i < 5; i++ {
				ch <- i
			}
		}()
		return ch
	}

	consumer := func(ch <-chan int) {
		for v := range ch {
			fmt.Printf("Recevied: %d\n", v)
		}
		fmt.Println("Done receiving!")
	}

	ch := owner()
	consumer(ch)

}
