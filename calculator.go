// Package calculator provides a library for simple calculations in Go.
package calculator

import (
	"fmt"
	"math"
)

// Add takes two numbers and returns the result of adding them together.
func Add(a, b float64) float64 {
	return a + b
}

// Subtract takes two numbers and returns the result of subtracting the second
// from the first.
func Subtract(a, b float64) float64 {
	return a - b
}

func Multiply(a, b float64) float64 {
	return a * b
}

// Divide takes two numbers and returns the result of dividing the first by the
// second, or an error if the second value is zero.
func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("bad input: %f, %f (division by zero is undefined)", a, b)
	}
	return a / b, nil
}

// Sqrt takes a number and returns its square root, or an error if the value is
// negative.
func SqRoot(a float64) (float64, error) {
	if a < 0 {
		return 0, fmt.Errorf("bad input: %f : (no negative numbers for Sqrt)", a)
	}
	return math.Sqrt(a), nil
}
