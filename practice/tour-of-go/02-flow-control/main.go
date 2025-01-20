package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func pow2(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim
}

func main() {
	// only has for statement, no parentheses
	{
		sum := 0
		for i := 0; i < 10; i++ {
			sum += i
		}
		fmt.Println(sum)
	}

	// init and post are optional
	{
		sum := 1
		for sum < 1000 {
			sum += sum
		}
		fmt.Println(sum)
	}

	// while loop can drop all semicolons
	{
		sum := 1
		for sum < 1000 {
			sum += sum
		}
		fmt.Println(sum)
	}

	// infinite loop
	//for {
	//
	//}

	// if statement has no parentheses, braces required
	{
		fmt.Println(sqrt(2), sqrt(-4))
	}

	// if with short statement
	{
		fmt.Println(
			pow(3, 2, 10),
			pow(3, 3, 20))
	}

	// if else
	{
		fmt.Println(
			pow2(3, 2, 10),
			pow2(3, 3, 20))
	}

	// switch
	{
		fmt.Print("Go runs on ")
		switch os := runtime.GOOS; os {
		case "darwin":
			fmt.Println("OS X.")
		case "linux":
			fmt.Println("Linux.")
		default:
			fmt.Printf("%s.\n", os)
		}
	}

	// switch evaluation order
	{
		fmt.Println("When's Saturday?")
		today := time.Now().Weekday()

		switch time.Saturday {
		case today + 0:
			fmt.Println("Today.")
		case today + 1:
			fmt.Println("Tomorrow.")
		case today + 2:
			fmt.Println("In two days.")
		default:
			fmt.Println("Too far away.")
		}
	}

	// switch w/ no condition, basically a long if-then-else chain
	{
		t := time.Now()
		switch {
		case t.Hour() < 12:
			fmt.Println("Good morning!")
		case t.Hour() < 17:
			fmt.Println("Good afternoon.")
		default:
			fmt.Println("Good evening.")
		}
	}

	// defer
	{
		defer fmt.Println("world")
		fmt.Println("hello")
	}

	// deferred calls are stored in a stack an executed LIFO
	{
		fmt.Println("counting")

		for i := 0; i < 10; i++ {
			defer fmt.Println(i)
		}

		fmt.Println("done")
	}
}
