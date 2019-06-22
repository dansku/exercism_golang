// Package bob provides all the tools for the bob exercism task.
package bob

import (
	"strings"
	"unicode"
)

// Hey gets a string and return the correct answer to the command
func Hey(remark string) string {

	// Clear remark
	remark = strings.Trim(remark, " \t\n\r")

	if isQuestion(remark) {
		if containLetter(remark) && isUpperCase(remark) {
			return "Calm down, I know what I'm doing!"
		}
		return "Sure."
	}

	if containLetter(remark) && isUpperCase(remark) {
		return "Whoa, chill out!"
	}

	if remark == "" {
		return "Fine. Be that way!"
	}

	return "Whatever."
}

// isQuestion check if the remark ends with ?
func isQuestion(remark string) bool {
	if strings.HasSuffix(remark, "?") {
		return true
	}
	return false
}

// isUpperCase checks if remark is uppercase
func isUpperCase(remark string) bool {
	if remark == strings.ToUpper(remark) {
		return true
	}
	return false
}

// containLetter returns true is remark contains letters
func containLetter(remark string) bool {
	for _, r := range remark {
		if unicode.IsLetter(r) {
			return true
		}
	}
	return false
}
