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
	ns   []float64
	want float64
}

func TestAdd(t *testing.T) {
	t.Parallel()

	testCases := []TestCase{
		{name: "Two identical numbers", ns: []float64{2, 2}, want: 4},
		{name: "One number is zero", ns: []float64{5, 0}, want: 5},
		{name: "More than two numbers", ns: []float64{1, 2, 3, 4}, want: 10},
		{name: "No numbers", ns: []float64{}, want: 0},
	}

	for _, tc := range testCases {
		got := calculator.Add(tc.ns...)
		if tc.want != got {
			t.Errorf("%s: Add(%f): want %f, got %f", tc.name, tc.ns, tc.want, got)
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
		{name: "Two identical numbers", ns: []float64{2, 0}, want: 2},
		{name: "One negative number", ns: []float64{1, 6}, want: -5},
		{name: "One decimal number", ns: []float64{2, 0.5}, want: 1.5},
		{name: "Multiple numbers", ns: []float64{2, 1, 1, 1}, want: -1},
		{name: "No numbers", ns: []float64{}, want: 0},
	}

	for _, tc := range testCases {
		got := calculator.Subtract(tc.ns...)
		if tc.want != got {
			t.Errorf("%s: Add(%f): want %f, got %f", tc.name, tc.ns, tc.want, got)
		}
	}
}

func TestMultiply(t *testing.T) {
	t.Parallel()

	testCases := []TestCase{
		{name: "Two positive numbers", ns: []float64{11, 9}, want: 99},
		{name: "Two negative numbers", ns: []float64{-2, -3.5}, want: 7},
		{name: "One negative number", ns: []float64{2, -3.5}, want: -7},
		{name: "Multiple numbers", ns: []float64{2, -3.5, -2}, want: 14},
		{name: "No numbers", ns: []float64{}, want: 0},
	}

	for _, tc := range testCases {
		got := calculator.Multiply(tc.ns...)
		if tc.want != got {
			t.Errorf("%s: Add(%f): want %f, got %f", tc.name, tc.ns, tc.want, got)
		}
	}
}

func TestDivide(t *testing.T) {
	t.Parallel()

	type divideTestCase struct {
		name        string
		ns          []float64
		want        float64
		errExpected bool
	}

	testCases := []divideTestCase{
		{name: "Regular division", ns: []float64{10, 5}, want: 2, errExpected: false},
		{name: "Decimal division", ns: []float64{10, 2.5}, want: 4, errExpected: false},
		{name: "Multiple numbers", ns: []float64{10, 2, 5}, want: 1, errExpected: false},
		{name: "Division by zero", ns: []float64{10, 0}, want: 0, errExpected: true},
		{name: "No numbers", ns: []float64{}, want: 0},
	}

	for _, tc := range testCases {
		got, err := calculator.Divide(tc.ns...)

		errReceived := err != nil

		if tc.errExpected != errReceived {
			t.Fatalf("%s: Divide(%f): %s", tc.name, tc.ns, err.Error())
		}

		if !tc.errExpected && tc.want != got {
			t.Errorf("%s: Divide(%f): want %f, got %f", tc.name, tc.ns, tc.want, got)
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

func TestEvaluate(t *testing.T) {
	t.Parallel()

	type TestCase struct {
		name        string
		expression  string
		want        float64
		errExpected bool
	}

	testCases := []TestCase{
		{name: "Multiplication", expression: "2 * 2", want: 4, errExpected: false},
		{name: "Addition", expression: "1 + 1.5", want: 2.5, errExpected: false},
		{name: "Division", expression: "18    /     6", want: 3, errExpected: false},
		{name: "Subtraction", expression: "100 - 0.1", want: 99.9, errExpected: false},
		{name: "Unknown operator", expression: "100 & 0.1", want: 0, errExpected: true},
		{name: "Prefix format", expression: "+ 1 2", want: 0, errExpected: true},
		{name: "Division by zero", expression: "10 / 0", want: 0, errExpected: true},
	}

	for _, tc := range testCases {
		got, err := calculator.Evaluate(tc.expression)

		errReceived := err != nil

		if tc.errExpected != errReceived {
			t.Fatalf("%s had unexpected error status: %s", tc.name, err.Error())
		}

		if !tc.errExpected && !closeEnough(tc.want, got, 0.001) {
			t.Errorf("%s: want %f, got %f", tc.name, tc.want, got)
		}
	}
}
