package main

import "testing"

type addTest struct {
	arg1, arg2, expected int
}

var addTests = []addTest{
	addTest{2, 5, 7},
	addTest{-2, 0, -2},
	addTest{11, 22, 33},
	addTest{1, 22, 23},
}

func TestAdd(t *testing.T) {

	for _, test := range addTests {
		got := Add(test.arg1, test.arg2)
		if got != test.expected {
			t.Errorf("Output %v not equal to expected %v", got, test.expected)
		}

	}
}

func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add(4, 6)
	}
}
