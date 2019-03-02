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

// Inequality determinesif the triangle is possible
func Inequality(a, b, c float64) bool {

	// If any side is zero, wrong triangle
	if a == 0 || b == 0 || c == 0 {
		return false
	}

	// Check if any side is infinite
	if math.IsInf(a, 0) == true || math.IsInf(b, 0) == true || math.IsInf(c, 0) == true {
		return false
	}

	//Get biggest value
	maxValue := math.Max(a, b)
	maxValue = math.Max(maxValue, c)

	if 2*maxValue <= a+b+c {
		return true
	}
	return false
}

// SameSides calculates how many equal sides a triangle has
func SameSides(a, c, b float64) int {

	if a == b && a == c {
		return 3
	}

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

	if sides == 1 {
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
	case 0:
		k = Sca
	case 2:
		k = Iso
	case 3:
		k = Equ
	}

	return k
}
