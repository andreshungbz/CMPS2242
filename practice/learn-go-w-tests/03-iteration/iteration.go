package iteration

import "strings"

// Takes a string and count and returns repeated string count times
func Repeat(character string, count int) string {
	var repeated strings.Builder

	for i := 0; i < count; i++ {
		repeated.WriteString(character)
	}

	return repeated.String()
}
