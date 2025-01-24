package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {
	persons := make([]Person, 0)

	for i := 0; i < 10_000_000; i++ {
		persons = append(persons, Person{"George", "Kalos", 24})
	}

	fmt.Println("Done")
}