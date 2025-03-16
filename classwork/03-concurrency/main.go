// Filename: main.go

package main

import (
	"fmt"
	"sync"
	"time"
)

func hello(wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(time.Second)
	fmt.Println("Hello")
}

func world(wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(time.Second)
	fmt.Println("World!")
}

func main() {
	// create a wait group
	var wg sync.WaitGroup
	
	// execute the Goroutine
	wg.Add(1)
	go hello(&wg)
	wg.Add(1)
	go world(&wg)

	wg.Wait()

	fmt.Println("Goodbye from main!")
}
