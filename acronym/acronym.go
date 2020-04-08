// Package acronym provides all the tools for the acronym exercism task.
package acronym

import (
	"regexp"
	"strings"
)

// Abbreviate gets a phrase and returns its acronyms
func Abbreviate(s string) string {
	var acronyms strings.Builder

	s = clearSentence(s)
	split := strings.Split(s, " ")

	for _, letter := range split {
		acronyms.WriteString(letter[0:1])
	}

	return strings.ToUpper(acronyms.String())
}

func clearSentence(s string) string {
	reg, _ := regexp.Compile("[^a-zA-Z0-9']+")
	return reg.ReplaceAllString(s, " ")
}
