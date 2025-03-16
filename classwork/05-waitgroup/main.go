package main

import (
	"fmt"
	"sync"
	"time"
)

func printName(wg *sync.WaitGroup, name string) {
	defer wg.Done()
	fmt.Println("Hello", name)
}

func printSum(wg *sync.WaitGroup, x, y int) {
	defer wg.Done()
	fmt.Printf("%d + %d = %d\n", x, y, x + y)
}

func waitSeconds(wg *sync.WaitGroup, seconds int) {
	defer wg.Done()
	time.Sleep(time.Second * time.Duration(seconds))
	fmt.Println("Done waiting")
}

func main() {
	var wg sync.WaitGroup
	wg.Add(3)

	go printName(&wg, "world")
	go printSum(&wg, 2, 4)
	go waitSeconds(&wg, 1)

	wg.Wait()

	fmt.Println("End of main")
}