package main

import "os"

func main() {
	panic("there is problem")

	_, err := os.Create("\tmp\file")

	if err != nil {
		panic(err)
	}
}
