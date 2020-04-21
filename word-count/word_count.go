// Package wordcount provides all the tools for the wordcount exercism task.
package wordcount

import (
	"regexp"
	"strings"
)

// Frequency is the map of word and it's occurencies
type Frequency map[string]int

// WordCount returns a map with the word frequency
func WordCount(phrase string) Frequency {

	frequency := make(map[string]int) // https://yourbasic.org/golang/gotcha-assignment-entry-nil-map/

	phrase = strings.ToLower(phrase)
	phrase = strings.Replace(phrase, ",", " ", -1)

	words := strings.Fields(phrase)

	reg := regexp.MustCompile("[^\\w']+")

	for _, word := range words {

		word = reg.ReplaceAllString(word, "")
		word = strings.Trim(word, "'")

		if _, found := frequency[word]; found {
			frequency[word]++
		} else {
			frequency[word] = 1
		}
	}

	return frequency
}
