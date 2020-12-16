// Package triangle is a utility for determining what kind of triangle
package triangle

import "math"

// Kind is the type of triangle
type Kind int

const (
	// NaT is not a triangle
	NaT Kind = iota
	// Equ is equilateral triangle
	Equ
	// Iso is isosceles triangle
	Iso
	// Sca is scalene triangle
	Sca
)

// KindFromSides should have a comment documenting it.
func KindFromSides(a, b, c float64) (k Kind) {
	// Write some code here to pass the test suite.
	// Then remove all the stock comments.
	// They're here to help you get started but they only clutter a finished solution.
	// If you leave them in, reviewers may protest!

	// Check if not a triangle
	if !isTriangle(a, b, c) {
		return NaT
	}

	// Assume scalene at first
	numSidesSame := 3

	// Subtract one for every equal side
	if a == b || a == c {
		numSidesSame--
	}
	if b == c {
		numSidesSame--
	}

	return Kind(numSidesSame)
}

func isTriangle(a, b, c float64) bool {
	// Any side is NaN
	if !isReal(a + b + c) {
		return false
	}

	// Any side has no length
	if a <= 0 || b <= 0 || c <= 0 {
		return false
	}

	// One side is way too long
	if a+b < c || a+c < b || b+c < a {
		return false
	}

	return true
}

// isReal checks if a float is a real number
func isReal(n float64) bool {
	if math.IsNaN(n) {
		return false
	}
	if math.IsInf(n, 1) {
		return false
	}
	if math.IsInf(n, -1) {
		return false
	}

	return true
}
