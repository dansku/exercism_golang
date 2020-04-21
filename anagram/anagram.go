package anagram

import (
	"strings"
)

// Detect gets a string and a string slice and detects the anagrams of it
func Detect(subject string, candidates []string) []string {

	var detectedAnagrams []string

	for _, word := range candidates {

		// Normalize test words
		var currentSubject = strings.ToLower(subject)
		var testWord = strings.ToLower(word)

		for _, letter := range testWord {
			if strings.Contains(currentSubject, string(letter)) && len(currentSubject) > 0 {
				currentSubject = strings.Replace(currentSubject, string(letter), "", 1)
			}
		}

		if len(currentSubject) == 0 && len(subject) == len(testWord) && strings.ToLower(subject) != testWord {
			detectedAnagrams = append(detectedAnagrams, word)
		}
	}

	return detectedAnagrams
}
