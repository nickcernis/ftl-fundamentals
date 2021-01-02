// Package calculator provides a library for simple calculations in Go.
package calculator

import (
	"errors"
	"math"
)

// Add takes numbers and returns the result of adding them together.
func Add(ns ...float64) float64 {
	result := 0.0
	for _, n := range ns {
		result += n
	}
	return result
}

// Subtract takes numbers and returns the result of subtracting them in turn.
func Subtract(ns ...float64) float64 {
	if len(ns) == 0 {
		return 0
	}

	result := ns[0]
	for _, n := range ns[1:] {
		result -= n
	}
	return result
}

// Multiply returns the product of multiple numbers.
func Multiply(ns ...float64) float64 {
	if len(ns) == 0 {
		return 0
	}

	result := ns[0]
	for _, n := range ns[1:] {
		result *= n
	}
	return result
}

// Divide returns numbers divided in turn or 0, error if division could not be performed.
func Divide(ns ...float64) (float64, error) {
	if len(ns) == 0 {
		return 0, nil
	}

	result := ns[0]
	for _, n := range ns[1:] {
		if n == 0 {
			return 0, errors.New("division by zero is undefined")
		}
		result /= n
	}

	return result, nil
}

// Sqrt returns the square root of x or 0, error if x is negative.
func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("the square root of a negative number is imaginary")
	}

	return math.Sqrt(x), nil
}
