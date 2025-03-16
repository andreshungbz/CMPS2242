package main

import (
	"fmt"
)

func main() {
	// create our first channel
	ch := make(chan int)

	// anonymous function
	go func() {
		ch <- 99
	}()

	// read from the channel
	x := <-ch
	fmt.Println(x)
}