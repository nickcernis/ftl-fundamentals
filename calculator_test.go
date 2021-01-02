package calculator_test

import (
	"calculator"
	"fmt"
	"math"
	"math/rand"
	"testing"
	"time"
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

func TestAddRandom(t *testing.T) {
	t.Parallel()

	rand.Seed(time.Now().UnixNano())
	n := 1

	for n <= 100 {
		rand1 := rand.Float64() * 1000
		rand2 := rand.Float64() * -1000

		want := rand1 + rand2
		got := calculator.Add(rand1, rand2)

		fmt.Printf("Random %f + %f = %f\n", rand1, rand2, got)

		if want != got {
			t.Errorf("Random add: Add(%f, %f): want %f, got %f", rand1, rand2, want, got)
		}

		n += 1
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

func closeEnough(a, b, tolerance float64) bool {
	return math.Abs(a-b) <= tolerance
}

func TestSqrt(t *testing.T) {
	t.Parallel()

	type sqrtTestCase struct {
		name        string
		x           float64
		want        float64
		errExpected bool
	}

	testCases := []sqrtTestCase{
		{name: "Regular square root", x: 10, want: 3.1622776602, errExpected: false},
		{name: "Decimal square root", x: 7.5, want: 2.7386127875, errExpected: false},
		{name: "Root of negative", x: -2, want: 0, errExpected: true},
	}

	for _, tc := range testCases {
		got, err := calculator.Sqrt(tc.x)

		errReceived := err != nil

		if tc.errExpected != errReceived {
			t.Fatalf("%s: Sqrt(%f): %s", tc.name, tc.x, err.Error())
		}

		if !tc.errExpected && !closeEnough(tc.want, got, 0.001) {
			t.Errorf("%s: Sqrt(%f): want %f, got %f", tc.name, tc.x, tc.want, got)
		}
	}
}
