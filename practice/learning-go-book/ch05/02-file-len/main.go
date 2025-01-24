package main

import (
	"errors"
	"fmt"
	"os"
)

func fileLen(filename string) (int, error) {
	if len(filename) == 0 {
		return 0, errors.New("empty string")
	}

	file, err := os.Open(filename)
	if err != nil {
		return 0, errors.New("file open error")
	}
	defer file.Close()

	fileInfo, _ := file.Stat()
	length := int(fileInfo.Size())

	return length, nil
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("one argument to filename needed")
		os.Exit(-1)
	}

	length, err := fileLen(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(-2)
	}

	fmt.Printf("Length of %v in bytes: %v\n", os.Args[1], length)
}