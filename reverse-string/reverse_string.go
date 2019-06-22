// Package reverse provides all the tools for the reverse exercism task.
package reverse

// Reverse gets a string and returns the same but reversed
func Reverse(str string) string {
	reverse := ""
	for _, r := range str {
		reverse = string(r) + reverse
	}
	return reverse
}
