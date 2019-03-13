// Package raindrops provides all the tools for the raindrops exercism task.
package raindrops

import "strconv"

// Convert gets an integer number and returns the right string output as requested in the task.
func Convert(number int) string {
	var raindrop string

	if number%3 == 0 {
		raindrop += "Pling"
	}

	if number%5 == 0 {
		raindrop += "Plang"
	}

	if number%7 == 0 {
		raindrop += "Plong"
	}

	if raindrop == "" {
		return strconv.Itoa(number)
	}

	return raindrop
}
