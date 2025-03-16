package main

import "fmt"

// specify my new type/container
// start with capital letter = public
// start with lowercase letter = private
type Employee struct {
	name     string
	salary   int
	currency string
}

// create methods on the Employee type
func (e Employee) displaySalary() {
	fmt.Printf("Salary of %s is %s%d", e.name, e.currency, e.salary)
}

// another method
func (e Employee) displayName() {
	fmt.Printf("This employee's name is %s", e.name)
}

func main() {
	john := Employee{
		name:     "John Smith",
		salary:   10000,
		currency: "BZ$",
	}

	jane := Employee{
		name:     "Jane Smith",
		salary:   15000,
		currency: "USD$",
	}

	fmt.Println(john.name, john.salary, john.currency)

	jane.displaySalary()
	fmt.Println()

	jane.displayName()
	fmt.Println()
}
