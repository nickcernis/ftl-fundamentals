package calculator_test

import (
	"calculator"
	"testing"
)

type TestCase struct {
	name string
	a, b float64
	want float64
}

func TestAdd(t *testing.T) {
	t.Parallel()

	testCases := []TestCase{
		{name: "Two identical numbers", a: 2, b: 2, want: 4},
		{name: "One number is zero", a: 5, b: 0, want: 5},
	}

	for _, tc := range testCases {
		got := calculator.Add(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("%s: Add(%f, %f): want %f, got %f", tc.name, tc.a, tc.b, tc.want, got)
		}
	}
}

func TestSubtract(t *testing.T) {
	t.Parallel()

	testCases := []TestCase{
		{name: "Two identical numbers", a: 2, b: 2, want: 0},
		{name: "One negative number", a: 1, b: 6, want: -5},
		{name: "One decimal number", a: 2, b: 0.5, want: 1.5},
	}

	for _, tc := range testCases {
		got := calculator.Subtract(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("%s: Subtract(%f, %f): want %f, got %f", tc.name, tc.a, tc.b, tc.want, got)
		}
	}
}

func TestMultiply(t *testing.T) {
	t.Parallel()

	testCases := []TestCase{
		{name: "Two positive numbers", a: 11, b: 9, want: 99},
		{name: "Two negative numbers", a: -2, b: -3.5, want: 7},
		{name: "One negative number", a: 2, b: -3.5, want: -7},
	}

	for _, tc := range testCases {
		got := calculator.Multiply(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("%s: Multiply(%f, %f): want %f, got %f", tc.name, tc.a, tc.b, tc.want, got)
		}
	}
}

func TestDivide(t *testing.T) {
	t.Parallel()

	type divideTestCase struct {
		name        string
		a, b        float64
		want        float64
		errExpected bool
	}

	testCases := []divideTestCase{
		{name: "Regular division", a: 10, b: 5, want: 2, errExpected: false},
		{name: "Decimal division", a: 10, b: 2.5, want: 4, errExpected: false},
		{name: "Division by zero", a: 10, b: 0, want: 0, errExpected: true},
	}

	for _, tc := range testCases {
		got, err := calculator.Divide(tc.a, tc.b)

		errReceived := err != nil

		if tc.errExpected != errReceived {
			t.Fatalf("%s: Divide(%f, %f): %s", tc.name, tc.a, tc.b, err.Error())
		}

		if !tc.errExpected && tc.want != got {
			t.Errorf("%s: Divide(%f, %f): want %f, got %f", tc.name, tc.a, tc.b, tc.want, got)
		}
	}
}
