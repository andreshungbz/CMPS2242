package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func MakePerson(firstName, lastName string, age int) Person {
	return Person{firstName, lastName, age}
}

func MakePersonPointer(firstName, lastName string, age int) *Person {
	return &Person{firstName, lastName, age}
}

func main() {
	person1 := MakePerson("George", "Kalos", 23)
	person2Pointer := MakePersonPointer("Sally", "Gates", 22)

	fmt.Println(person1)
	fmt.Println(*person2Pointer)
}
