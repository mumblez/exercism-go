// Package triangle calculates the type of triangle given the lengths
package triangle

import "math"

// Kind type
type Kind string

const (
	testVersion = 3
	NaT         = Kind("NaT") // not a triangle
	Equ         = Kind("Equ") // equilateral
	Iso         = Kind("Iso") // isoceles
	Sca         = Kind("Sca") // scalene
)

// KindFromSides given length of all sides of triangle will return the type of triangle
func KindFromSides(a, b, c float64) Kind {
	// NaT check 1
	if a <= 0 || b <= 0 || c <= 0 ||
		math.IsNaN(a) || math.IsNaN(b) || math.IsNaN(c) ||
		math.IsInf(a, 0) || math.IsInf(b, 0) || math.IsInf(c, 0) {
		return NaT
	}

	// Nat Check 2
	if !((a + b) < c) && !((a + c) < b) && !((b + c) < a) {
		switch {
		// Equality check
		case (a == b && a == c):
			return Equ
		// Isoceles Check
		case (a == b || a == c || b == c):
			return Iso
		// Scalene Check
		case a != b && a != c && b != c:
			return Sca
		}
	}
	return NaT
}
