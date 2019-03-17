// Package scrabble provides all the tools for the scrabble exercism task.
package scrabble

import "unicode"

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
