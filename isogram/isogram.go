// Package isogram provides all the tools for the isogram exercism task.
package isogram

import (
	"strings"
	"unicode"
)

// IsIsogram gets a word or phrase and returns if isograms or not
func IsIsogram(word string) bool {

	replacer := strings.NewReplacer(" ", "", "-", "")
	word = replacer.Replace(word)
	chars := map[rune]bool{}

	for _, l := range word {
		l = unicode.ToLower(l)
		if chars[l] {
			return false
		}
		chars[l] = true
	}

	return true
}
