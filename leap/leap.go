// Package leap provides all the tools for the leap exercism task.
package leap

// IsLeapYear checks if year entered is true or not
func IsLeapYear(year int) bool {
	return (year%100 != 0 || year%400 == 0) && year%4 == 0
}
