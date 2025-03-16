package main

import "fmt"

func makeAdders() (func(int) int, func(int) int) {
	sum := 0

	inc := func(x int) int {
		sum += x
		return sum
	}

	double := func(x int) int {
		sum += x * 2
		return sum
	}

	return inc, double
}

// all variables on stack
func normalFunc(x int) int {
	y := x + 1
	return y
}

// closure
func makeClosure() func() int {
	x := 1
	return func() int {
		return x
	}
}

// closure
func counter() func() int {
	n := 0
	return func() int {
		n++
		return n
	}
}

// func main() {
// 	inc, double := makeAdders()

// 	fmt.Println(inc(1))
// 	fmt.Println(double(2))
// 	fmt.Println(inc(3))
// }

// func main() {
// 	fmt.Println(normalFunc(2))
// }

// func main() {
// 	f := makeClosure()
// 	fmt.Println(f())
// }

// func main() {
// 	x := 1
// 	func() {
// 		fmt.Println(x)
// 	}()
// }

func main() {
	f := counter()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
}