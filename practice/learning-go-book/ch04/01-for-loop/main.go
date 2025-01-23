package main

import (
	"fmt"
	"math/rand"
)

func main() {
	randomNumbers := make([]int, 0, 100)

	for i := 0; i < 100; i++ {
		randomNumbers = append(randomNumbers, rand.Intn(100))
	}

	fmt.Println(randomNumbers)

	for _, number := range randomNumbers {
		switch {
		case number % 2 == 0:
			fmt.Println("Two!")
		case number % 3 == 0:
			fmt.Println("Three!")
		case number % 2 == 0 && number % 3 == 0:
			fmt.Println("Six!")
		default:
			fmt.Println("Never mind.")
		}
	}
}