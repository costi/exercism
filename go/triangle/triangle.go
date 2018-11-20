// Package triangle does math about triangles
package triangle

import "math"

// Kind is what type of triangle: equilateral, isosceles, scalene
type Kind string

const (
	NaT = "Not a Triangle" // not a triangle
	Equ = "Equilateral"    // equilateral
	Iso = "Isosceles"      // isosceles
	Sca = "Scalene"        // scalene
)

// KindFromSides returns the type of triangle: eq, iso, scalene
func KindFromSides(a, b, c float64) Kind {
	var k Kind
	if !isTriangle(a, b, c) {
		return NaT
	}

	if a == b && b == c {
		k = Equ
	} else if a == b || a == c || b == c {
		k = Iso
	} else {
		k = Sca
	}
	return k
}

func isTriangle(a, b, c float64) bool {
	inf := math.Inf(1)
	if a > 0 && b > 0 && c > 0 && a+b >= c && a+c >= b && b+c >= a && a != inf && b != inf && c != inf {
		return true
	}
	return false
}
