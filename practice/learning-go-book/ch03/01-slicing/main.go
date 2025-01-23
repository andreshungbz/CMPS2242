package main

import "fmt"

func main() {
	greetings := []string{"Hello", "Hola", "नमस्कार", "こんにちは", "Привіт"}
	firstTwo := greetings[:2]
	secondThirdFourth := greetings[1:4]
	fourthFifth := greetings[3:]

	fmt.Println(greetings)
	fmt.Println(firstTwo)
	fmt.Println(secondThirdFourth)
	fmt.Println(fourthFifth)
}