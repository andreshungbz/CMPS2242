package main

import "fmt"

func UpdateSlice(slice []string, s string) {
	slice[len(slice) - 1] = s
	fmt.Println("inside UpdateSlice:", slice)
}

func GrowSlice(slice []string, s string) {
	slice = append(slice, s)
	fmt.Println("inside GrowSlice:", slice)
}

func main() {
	names := []string{"George", "Sally", "Mallory"}

	fmt.Println(names)
	UpdateSlice(names, "Poppy")
	fmt.Println(names)

	fmt.Print("\n")

	fmt.Println(names)
	GrowSlice(names, "Larry")
	fmt.Println(names)
}