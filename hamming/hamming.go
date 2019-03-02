// Package hamming provides all the tools for the hamming exercism task.
package hamming

import "errors"

// Distance calculates the hamming distance between two DNA strands, returning the distance value.
func Distance(a, b string) (int, error) {

	// Check if a and b has the same lenght
	if len(a) != len(b) {
		return 0, errors.New("DNA strands must have the same lenght")
	}

	// Iterate between the two DNA strands and calculate hamming distance
	hammingDistance := 0
	for i := range a {
		if a[i] != b[i] {
			hammingDistance++
		}
	}
	return hammingDistance, nil
}
