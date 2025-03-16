// main.go

package main

import (
	"fmt"
	"strconv"
)

func promptInput(msg string) string {
	var result string
	fmt.Print(msg + ": ")
	fmt.Scanln(&result)
	return result
}

func main() {
	firstInput := promptInput("Enter first number")
	firstNumber, err := strconv.ParseInt(firstInput, 10, 0)
	if err != nil {
		fmt.Println("Invalid first number:", err)
		return
	}

	secondInput := promptInput("Enter second number")
	secondNumber, err := strconv.ParseInt(secondInput, 10, 0)
	if err != nil {
		fmt.Println("Invalid second number:", err)
		return
	}

	sum := firstNumber + secondNumber
	fmt.Printf("Sum: %d\n", sum)
}
