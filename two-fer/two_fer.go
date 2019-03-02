// Package twofer provides all the tools for the twofer exercism task.
package twofer

import "fmt"

// ShareWith checks if function contains name and returns string accordingly.
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return fmt.Sprintf("One for %s, one for me.", name)
}
