// Package triangle determine if a triangle is equilateral, isosceles, or scalene.
package triangle

import "math"

type Kind int

const (
	NaT Kind = iota // not a triangle
	Equ             // equilateral
	Iso             // isosceles
	Sca             // scalene
)

// KindFromSides determines if a triangle is equilateral, isosceles, or scalene.
func KindFromSides(a, b, c float64) Kind {
	var k Kind
	if math.IsNaN(a) || math.IsInf(a, 0) || math.IsNaN(b) || math.IsInf(b, 0) || math.IsNaN(c) || math.IsInf(c, 0) {
		k = NaT
	} else if a <= 0 || b <= 0 || c <= 0 || a+b < c || b+c < a || a+c < b {
		k = NaT
	} else if a == b && a == c {
		k = Equ
	} else if a == b || a == c || b == c {
		k = Iso
	} else {
		k = Sca
	}
	return k
}
