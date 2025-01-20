package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"math/rand"
)

// C-style print functions
func random() {
	fmt.Println("My favorite number is", rand.Intn(10))
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
	fmt.Println(math.Pi)
}

// datatype goes after
func add(x, y int) int {
	return x + y
}

// return multiple values
func swap(x, y string) (string, string) {
	return y, x
}

// return named values; variables declared @ top of function
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return // return with no arguments returns named values, only use in short functions
}

// declare a variable; types goes last
var c, python, java bool // package level

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func needInt(x int) int {
	return x*10 + 1
}

func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	random()
	fmt.Println(add(42, 13))
	a, b := swap("Hello", "World")
	fmt.Println(a, b)
	fmt.Println(split(17))

	// function level
	{
		var i int
		fmt.Println(i, c, python, java)
	}

	// variable initializer don't need type
	{
		var i, j int = 1, 2
		var k, l, m = true, false, "no!"
		fmt.Println(i, j, k, l, m)
	}

	// short assignment operator
	{
		var i, j int = 1, 2
		k := 3
		c, python, java := true, false, "no!"

		fmt.Println(i, j, k, c, python, java)
	}

	// basic types
	{
		fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
		fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
		fmt.Printf("Type: %T Value: %v\n", z, z)
	}

	// default initial value is zero
	{
		var i int
		var f float64
		var b bool
		var s string
		fmt.Printf("%v %v %v %q\n", i, f, b, s)
	}

	// type conversions
	{
		var x, y int = 3, 4
		var f float64 = math.Sqrt(float64(x*x + y*y))
		var z uint = uint(f)
		fmt.Println(x, y, z)
	}

	// type inference
	{
		v := 42.5
		fmt.Printf("v is of type %T\n", v)
	}

	// constants
	{
		const Pi = 3.14
		const World = "Earth"
		fmt.Println("Hello", World)
		fmt.Println("Happy", Pi, "Day")
		const Truth = true
		fmt.Println("Go rules?", Truth)
	}

	// numeric constants
	{
		const (
			Big   = 1 << 100
			Small = Big >> 99
		)

		fmt.Println(needInt(Small))
		fmt.Println(needFloat(Small))
		fmt.Println(needFloat(Big))
	}
}
