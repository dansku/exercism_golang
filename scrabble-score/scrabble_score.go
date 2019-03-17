// Package scrabble provides all the tools for the scrabble exercism task.
package scrabble

import "unicode"

// Benchmark results with code improvements
//
// Initially with map inside function:
// BenchmarkScore-8          100000             16328 ns/op            6384 B/op         45 allocs/op
//
// Moving map outside function:
// BenchmarkScore-8         1000000              1106 ns/op              32 B/op          4 allocs/op
//
// Changing strings.ToLower to unicode.ToLower
// BenchmarkScore-8         2000000               985 ns/op               0 B/op          0 allocs/op
//
// Using switch instead of map
// BenchmarkScore-8         5000000               327 ns/op               0 B/op          0 allocs/op

var points = map[rune]int{
	'a': 1, 'e': 1, 'i': 1, 'o': 1, 'u': 1, 'l': 1, 'n': 1, 'r': 1, 's': 1, 't': 1,
	'd': 2, 'g': 2,
	'b': 3, 'c': 3, 'm': 3, 'p': 3,
	'f': 4, 'h': 4, 'v': 4, 'w': 4, 'y': 4,
	'k': 5,
	'j': 8, 'x': 8,
	'q': 10, 'z': 10,
}

// getScores returns the scrabble value for a given letter
func getScores(letter rune) int {
	switch letter {
	case 'a', 'e', 'i', 'o', 'u', 'l', 'n', 'r', 's', 't':
		return 1
	case 'd', 'g':
		return 2
	case 'b', 'c', 'm', 'p':
		return 3
	case 'f', 'h', 'v', 'w', 'y':
		return 4
	case 'k':
		return 5
	case 'j', 'x':
		return 8
	case 'q', 'z':
		return 10
	default:
		return 0
	}
}

// Score calculates and returns the scrabble score of a given word
func Score(word string) int {
	score := 0
	for _, letter := range word {
		letter = unicode.ToLower(letter)
		score += getScores(letter)
	}
	return score
}
