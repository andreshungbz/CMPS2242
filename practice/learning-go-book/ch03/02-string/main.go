package main

import "fmt"

func main() {
	message := "Hi 👀 and 💙"
	runes := []rune(message)
	fmt.Println(string(runes[3]))
	fmt.Println(runes)
	fmt.Println([]byte(message))
}