package main

import (
	"fmt"
	"sync"
)

func worker(wg *sync.WaitGroup, id int) {
	defer wg.Done()
	fmt.Printf("Worker %d started\n", id)
	// time.Sleep(2 * time.Second) // Replace this
	fmt.Printf("Worker %d finished\n", id)
}

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(&wg, i)
	}

	// time.Sleep(3 * time.Second) // Replace this
	wg.Wait()

	fmt.Println("All workers done")
}
