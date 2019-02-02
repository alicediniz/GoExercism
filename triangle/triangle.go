// Package triangle provides a "detection triangle" function.
package triangle

import "math"

// Kind is an integer type used to hold a value between 0 to 2
// which indicates a shape
type Kind int

const (
	// NaT : not a triangle
	NaT Kind = iota
	// Equ : equilateral
	Equ
	// Iso : isosceles
	Iso
	// Sca : scalene
	Sca
)

// KindFromSides receives 3 sides of a geometric figure and informs what kind of shape it is,
//which can be a triangle (equilateral, isosceles or scalene) or not.
func KindFromSides(a, b, c float64) Kind {
	verifyType := (math.IsNaN(a) || math.IsNaN(b) || math.IsNaN(c))
	verifyGreaterZero := (a <= 0 || b <= 0 || c <= 0)
	verifySums := ((a+b < c) || (b+c < a) || (a+c < b))
	verifyInf := (math.IsInf(a, 0) || math.IsInf(b, 0) || math.IsInf(c, 0))

	if verifyType || verifyGreaterZero || verifySums || verifyInf {
		return NaT
	}
	if (a == b) && (b == c) {
		return Equ
	}
	if (a == b) || (b == c) || (a == c) {
		return Iso
	}
	return Sca
}
