// Package proverb provides all the tools for the proverb exercism task.
package proverb

import "fmt"

// Proverb takes a slice or words and returns the proverb.
func Proverb(rhyme []string) []string {
	var finalRhyme []string

	for i := range rhyme {
		if i < len(rhyme)-1 {
			finalRhyme = append(finalRhyme, fmt.Sprintf("For want of a %s the %s was lost.", rhyme[i], rhyme[i+1]))
		} else {
			finalRhyme = append(finalRhyme, fmt.Sprintf("And all for the want of a %s.", rhyme[0]))
		}
	}
	return finalRhyme
}
