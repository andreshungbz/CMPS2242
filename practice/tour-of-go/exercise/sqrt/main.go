package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	// initial value
	result := 1.0

	// calculate the next value and check if it is the same as the previous
	// if so, the square root is found and loop can be exited.
	// else set result as next, print feedback, and continue loop

	for i := 1; ; i++ { // have a counter to print iteration number
		next := result - (result*result-x)/(2*result) // Newton's method

		if (math.Abs(next - result)) == 0 { // the absolute value is taken for negative cases
			break
		}

		result = next
		fmt.Printf("[iteration %v] %v\n", i, result)
	}

	return result
}

func main() {
	input := 2.0
	result := Sqrt(input)
	fmt.Printf("\nSqrt(%v) answer: %v\n", input, result)
	fmt.Printf("math.Sqrt(%v) answer: %v\n", input, math.Sqrt(input))
}
