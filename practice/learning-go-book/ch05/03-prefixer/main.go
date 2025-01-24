package main

import "fmt"

func prefixer(prefix string) func(s string) string {
	return func(s string) string {
		return prefix + " " + s
	}
}

func main() {
	helloPrefix := prefixer("Hello")
	fmt.Println(helloPrefix("Bob"))
	fmt.Println(helloPrefix("Maria"))
}