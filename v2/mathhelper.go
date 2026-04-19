// Package mathhelper provides simple mathematical helper functions.
package mathhelper

import "golang.org/x/exp/constraints"

// Number is any integer or floating-point type.
type Number interface {
	constraints.Integer | constraints.Float
}

// Add returns the sum of a and b.
//
// For more on addition, see
// https://www.mathsisfun.com/numbers/addition.html.
func Add[T Number](a, b T) T {
	return a + b
}
