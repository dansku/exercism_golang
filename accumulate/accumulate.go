// Package accumulate provides all the tools for the accumulate exercism task.
package accumulate

// Accumulate gets the string and a response and convert the text to what is needed
func Accumulate(input []string, converter func(string) string) []string {

	var acc []string

	for _, word := range input {
		acc = append(acc, converter(word))
	}
	return acc

}
