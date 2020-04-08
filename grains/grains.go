// Package grains provides all the tools for the grain exercism task.
package grains

import (
	"errors"
)

// Square calulates the total amount of grains of wheat on a chessboard square
func Square(squares int) (uint64, error) {
	if squares < 1 || squares > 64 {
		return 0, errors.New("Error found")
	}
	// 1<<n = 2^n
	return 1 << uint(squares-1), nil
}

// Total calculates the total amount of wheat grains on a chessboard
func Total() uint64 {
	return 1<<64 - 1
}
