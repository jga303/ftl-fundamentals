package calculator_test

import (
	"calculator"
	"math/rand"
	"testing"
	"time"
)

type testCase struct {
	a, b float64
	want float64
	desc string
}

type testCaseWithErr struct {
	a, b        float64
	want        float64
	desc        string
	errExpected bool
}

func TestAdd(t *testing.T) {
	t.Parallel()

	testCases := []testCase{
		{a: 2, b: 5, want: 7, desc: "Sum two positives"},
		{a: -6, b: 2, want: -4, desc: "Sum a negative and a positive"},
		{a: -5, b: -20, want: -25, desc: "Sum two negatives"},
		{a: 2.5, b: -1.5, want: 1, desc: "Sum fractions"},
	}

	for _, tc := range testCases {
		got := calculator.Add(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("Add(%.2f, %.2f): want %.2f, got %.2f (%s)", tc.a, tc.b, tc.want, got, tc.desc)
		}
	}
}

// a test function with generated random numbers:
func TestAddRandom(t *testing.T) {
	t.Parallel()
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 100; i++ {
		a := rand.Float64()
		b := rand.Float64()
		want := a + b
		got := calculator.Add(a, b)
		if want != got {
			t.Fatalf("Add(%f, %f): want %f, got %f", a, b, want, got)
		}
	}
}

func TestSubtract(t *testing.T) {
	t.Parallel()

	testCases := []testCase{
		{a: 2, b: 5, want: -3, desc: "Subtract two positives"},
		{a: -6, b: 8, want: -14, desc: "Subtract a negative and a positive"},
		{a: -5, b: -20, want: 15, desc: "Subtract two negatives"},
		{a: 2.3, b: -1.5, want: 3.8, desc: "Subtract fractions"},
	}

	for _, tc := range testCases {
		got := calculator.Subtract(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("Subtract(%.2f, %.2f): want %.2f, got %.2f (%s)", tc.a, tc.b, tc.want, got, tc.desc)
		}
	}
}

func TestMultiply(t *testing.T) {
	t.Parallel()

	testCases := []testCase{
		{a: 3, b: 0, want: 0, desc: "Multiply two positives"},
		{a: -3, b: +2, want: -6, desc: "Multiply a negative and a positive"},
		{a: -5, b: -2, want: 10, desc: "Multiply two negatives"},
		{a: 2.5, b: -1.5, want: -3.75, desc: "Multiply fractions"},
	}

	for _, tc := range testCases {
		got := calculator.Multiply(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("Multiply(%.2f, %.2f): want %.2f, got %.2f (%s)", tc.a, tc.b, tc.want, got, tc.desc)
		}
	}
}

func TestDivide(t *testing.T) {
	t.Parallel()

	testCases := []testCaseWithErr{
		//fails because precision of float !: {a: 1, b: 3, want: 0.33, desc: "Divide two positives", errExpected: false},
		{a: 9, b: 3, want: 3, desc: "Divide two positives", errExpected: false},
		{a: -7, b: +2, want: -3.5, desc: "Divide a negative and a positive", errExpected: false},
		{a: -5, b: -2, want: 2.5, desc: "Divide two negatives", errExpected: false},
		{a: 4.5, b: -1.5, want: -3, desc: "Divide fractions", errExpected: false},
		{a: 3, b: 0, want: 999, desc: "Divide by zero", errExpected: true},
	}

	for _, tc := range testCases {
		got, err := calculator.Divide(tc.a, tc.b)
		errReturned := err != nil

		if errReturned != tc.errExpected {
			t.Fatalf("Divide(%.2f, %.2f): unexpected error status: %v (%s)", tc.a, tc.b, err, tc.desc)
		}

		if !tc.errExpected && tc.want != got {
			t.Errorf("Divide(%.2f, %.2f): want %.2f, got %.2f (%s)", tc.a, tc.b, tc.want, got, tc.desc)
		}
	}
}

func TestSqRoot(t *testing.T) {
	t.Parallel()

	testCases := []testCaseWithErr{
		{a: 9, want: 3, desc: "SqRoot positives", errExpected: false},
		{a: 1000000, want: 1000, desc: "SqRoot big", errExpected: false},
		//{a: 2, want: 1.41, desc: "SqRoot big", errExpected: false},
		{a: 0, want: 0, desc: "SqRoot zero", errExpected: false},
		{a: 2.25, want: 1.5, desc: "SqRoot fractions", errExpected: false},
		{a: -1, want: 999, desc: "SqRoot of negatives undefined", errExpected: true},
	}

	for _, tc := range testCases {
		got, err := calculator.SqRoot(tc.a)
		errReturned := err != nil

		if errReturned != tc.errExpected {
			t.Fatalf("SqRoot(%.2f): unexpected error status: %v (%s)", tc.a, err, tc.desc)
		}

		if !tc.errExpected && tc.want != got {
			t.Errorf("SqRoot(%.2f): want %.2f, got %.2f (%s)", tc.a, tc.want, got, tc.desc)
		}
	}
}
