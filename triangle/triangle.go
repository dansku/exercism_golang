// Package triangle provides all the tools for the twofer exercism task.
package triangle

import (
	"math"
)

// Kind is the type of triengle
type Kind int

const (
	// NaT Value if not a triangle
	NaT = -1

	// Equ for equilateral triangles
	Equ = 3

	// Iso for isosceles triangles
	Iso = 2

	// Sca for scalene triangles
	Sca = 0
)

// ZeroOrInfinite returns true if number is zero or infinite.
func ZeroOrInfinite(a float64) bool {
	if a == 0 || math.IsInf(a, 0) == true {
		return true
	}
	return false
}

// Inequality determinesif the triangle is possible.
func Inequality(a, b, c float64) bool {

	// If any side is zero, wrong triangle
	if ZeroOrInfinite(a) || ZeroOrInfinite(b) || ZeroOrInfinite(c) {
		return false
	}

	//Get max value
	maxValue := math.Max(a, math.Max(b, c))

	if 2*maxValue <= a+b+c {
		return true
	}
	return false
}

// SameSides calculates how many equal sides a triangle has
func SameSides(a, c, b float64) int {

	sides := 0

	if a == b {
		sides++
	}

	if a == c {
		sides++
	}

	if b == c {
		sides++
	}

	return sides
}

// KindFromSides gets three values which will be used to define if it can create a triangle and which kind of triangle it is.
func KindFromSides(a, b, c float64) Kind {

	// Test triangle inequality
	if !Inequality(a, b, c) {
		return NaT
	}

	triangleSameSides := SameSides(a, b, c)

	var k Kind
	switch triangleSameSides {
	case 3:
		k = Equ
	case 1:
		k = Iso
	default:
		k = Sca
	}

	return k
}
