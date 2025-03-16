// Filename: hello.go

package main

import (
	"fmt"
	"strconv"
)

func main() {
	// ask for a quantity
	var quantityInput string
	fmt.Print("Enter the quantity of items: ")
	fmt.Scanln(&quantityInput)

	// try to convert input into an integer value
	quantity, err := strconv.ParseInt(quantityInput, 10, 0)
	if err != nil {
		fmt.Println("Invalid quantity: ", err)
		return
	}

	// ask for a price
	var priceInput string
	fmt.Print("Enter the price per item: $")
	fmt.Scanln(&priceInput)

	// try to convert input into a float value
	price, err := strconv.ParseFloat(priceInput, 64)
	if err != nil {
		fmt.Println("Invalid price: ", err)
		return
	}

	// print total
	total := float64(quantity) * price
	fmt.Printf("Total: $%.2f\n", total)
}
