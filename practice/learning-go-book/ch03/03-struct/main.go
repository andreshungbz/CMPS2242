package main

import "fmt"

func main() {
	type Employee struct {
		firstName string
		lastName string
		id int
	}

	first := Employee {"George", "Kalos", 1}

	second := Employee {
		firstName: "Bill",
		lastName: "Thompson",
		id: 2,
	}

	var third Employee
	third.firstName = "Sally"
	third.lastName = "Gates"
	third.id = 3

	fmt.Println(first)
	fmt.Println(second)
	fmt.Println(third)
}