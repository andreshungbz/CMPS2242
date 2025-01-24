package main

import (
	"errors"
	"fmt"
	"strconv"
)

func add(i, j int) (int, error) { return i + j, nil }
func sub(i, j int) (int, error) { return i - j, nil }
func mul(i, j int) (int, error) { return i * j, nil }
func div(i, j int) (int, error) {
  result := 0
  if j == 0 {
    return result, errors.New("division by zero")
  }
  return i / j, nil
}

var operationMap = map[string]func(int, int) (int, error){
	"+": add,
	"-": sub,
	"*": mul,
	"/": div,
}

func main() {
	expressions := [][]string{
		{"2", "+", "3"},
		{"2", "-", "3"},
		{"2", "*", "3"},
		{"2", "/", "3"},
		{"two", "+", "three"},
		{"5"},
    {"2", "/", "0"},
	}

	for _, expression := range expressions {
		if len(expression) != 3 {
			fmt.Println("invalid expression:", expression)
			continue
		}

		p1, err := strconv.Atoi(expression[0])
		if err != nil {
			fmt.Println(err)
			continue
		}

		operator := expression[1]
		opFunction, ok := operationMap[operator]
		if !ok {
			fmt.Println("unsupported operator:", operator)
      continue
		}

		p2, err := strconv.Atoi(expression[2])
		if err != nil {
			fmt.Println(err)
			continue
		}

		result, err := opFunction(p1, p2)
    if err != nil {
      fmt.Println(err)
      continue
    }

		fmt.Println(result)
	}
}
