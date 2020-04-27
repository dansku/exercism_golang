package pangram

import "fmt"

// IsPangram checks if string contains all letters in the alphabet
func IsPangram(input string) bool {

	for i := 'a'; i <= 'z'; i++ {
		fmt.Printf("%v", i)
	}
	return false
}

// InSlice takes a slice and looks for an element in it
func InSlice(slice []rune, val rune) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}
