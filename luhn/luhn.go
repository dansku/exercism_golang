// Package luhn provides all the tools for the luhn exercism task.
package luhn

import (
	"strconv"
	"strings"
)

// Reverse string order so that the calculations are from right to left
func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

// Valid checks if a string is valid in the luhn algorithm
func Valid(input string) bool {

	// Remove empty spaces
	input = strings.Replace(input, " ", "", -1)

	// Check string lenght
	if len(input) <= 1 {
		return false
	}

	// Check if only numbers
	_, err := strconv.Atoi(input)
	if err != nil {
		return false
	}

	final := 0
	for i, num := range Reverse(input) {
		number := int(num) - 48
		if i%2 != 0 {
			number = number * 2
			if number > 9 {
				number = number - 9
			}
		}
		final += number
	}

	if final%10 != 0 {
		return false
	}

	return true
}
