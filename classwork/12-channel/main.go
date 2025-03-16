package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, ch chan int) {
	for task := range ch {
		fmt.Printf("Worker %d processing task %d\n", id, task)
		time.Sleep(time.Second * 1)
	}
}

func main() {
	tasks := make(chan int, 5)
	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			worker(i, tasks)
		}(i)
	}

	for j := 1; j <= 10; j++ {
		tasks <- j
		fmt.Println("Main sent to tasks channel:", j)
	}
	close(tasks)

	wg.Wait()
}
