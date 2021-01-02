// Package calculator provides a library for simple calculations in Go.
package calculator

import (
	"errors"
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

// Multiply returns the product of two numbers.
func Multiply(a, b float64) float64 {
	return a * b
}

// Divide returns a divided by b or 0, error if division could not be performed.
func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero is undefined")
	}

	return a / b, nil
}

// Sqrt returns the square root of x or 0, error if x is negative.
func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("the square root of a negative number is imaginary")
	}

	return math.Sqrt(x), nil
}
