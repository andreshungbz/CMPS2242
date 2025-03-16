package main

import (
	"fmt"
	"strconv"
)

// interface

type Shape interface {
	Area() float64
}

// structs

type Circle struct {
	Radius float64
}

type Rectangle struct {
	Length float64
	Width  float64
}

type Square struct {
	Length float64
}

type Triangle struct {
	Base   float64
	Height float64
}

type Parallelogram struct {
	Base   float64
	Height float64
}

// struct methods

func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

func (r Rectangle) Area() float64 {
	return r.Length * r.Width
}

func (s Square) Area() float64 {
	return s.Length * s.Length
}

func (t Triangle) Area() float64 {
	return t.Base * t.Height * 0.5
}

func (p Parallelogram) Area() float64 {
	return p.Base * p.Height
}

// functions

func PrintArea(s Shape) {
	fmt.Printf("Area: %.2f\n", s.Area())
}

// func PrintCircleArea(c Circle) {
// 	fmt.Printf("Area: %.2f\n", c.Area())
// }

// func PrintRectangleArea(r Rectangle) {
// 	fmt.Printf("Area: %.2f\n", r.Area())
// }

// func PrintSquareArea(s Square) {
// 	fmt.Printf("Area: %.2f\n", s.Area())
// }

// func PrintTriangleArea(t Triangle) {
// 	fmt.Printf("Area: %.2f\n", t.Area())
// }

// func PrintParallelogramArea(p Parallelogram) {
// 	fmt.Printf("Area: %.2f\n", p.Area())
// }

func main() {

	{
		circle := Circle{}
		fmt.Println("[Circle]")

		fmt.Print("Enter Radius: ")
		var input string
		fmt.Scanln(&input)

		radius, err := strconv.ParseFloat(input, 64)
		if !checkError(err) {
			circle.Radius = radius
			PrintArea(circle)
		}
	}

	{
		rectangle := Rectangle{}
		fmt.Println("[Rectangle]")

		fmt.Print("Enter Length: ")
		var input string
		fmt.Scanln(&input)

		fmt.Print("Enter Width: ")
		var input2 string
		fmt.Scanln(&input2)

		length, err := strconv.ParseFloat(input, 64)
		width, err2 := strconv.ParseFloat(input2, 64)

		if !checkError(err) || !checkError(err2) {
			rectangle.Length = length
			rectangle.Width = width
			PrintArea(rectangle)
		}
	}

	// rectangle := Rectangle{}
	// square := Square{}
	// triangle := Triangle{}
	// parallelogram := Parallelogram{}

	// shapes := []Shape{circle, rectangle, square, triangle, parallelogram}

}

func checkError(e error) bool {
	if e != nil {
		fmt.Println("Invalid input! Skipping...")
		return true
	}

	return false
}
