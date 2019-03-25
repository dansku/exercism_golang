// Package diffsquares provides all the tools for the diffsquares exercism task.
package diffsquares

//SquareOfSum gets a natural number and returns the square of sum of numbers from 1..n
func SquareOfSum(n int) int {
	total := 0
	for i := 1; i <= n; i++ {
		total += i
	}
	return total * total
}

// SumOfSquares gets a natural number and returns the sum of square numbers from 1..n
func SumOfSquares(n int) int {
	total := 0
	for i := 1; i <= n; i++ {
		total += i * i
	}
	return total
}

// Difference gets a natural number and returns the difference between square of the sums and the sum of quares
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
