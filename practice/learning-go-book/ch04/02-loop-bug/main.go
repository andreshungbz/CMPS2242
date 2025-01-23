package main

import "fmt"

func main() {
	var total int

	for i := 0; i < 10; i++ {
		total := total + i // BUG: := creates a new variable that shadows outer total, fix using = instead
		fmt.Println(total)
	}

	fmt.Printf("Total: %v\n", total);
}