package basics

import "testing"

func TestEuclidianDiv(t *testing.T) {
	type input_struct struct {
		d1, d2, expected_q, expected_r int
	}
	inputs := [2]input_struct{
		{10, 2, 5, 0},
		{11, 5, 2, 1},
	}
	for _, input := range inputs {
		q, r := EuclidianDiv(input.d1, input.d2)
		if q != input.expected_q || r != input.expected_r {
			t.Errorf("got: %d, %d; want: %d, %d", q, r, input.expected_q, input.expected_r)
		}
	}
}

func TestMin3(t *testing.T) {
	type input_struct struct {
		a, b, c, expected int
	}
	inputs := [3]input_struct{
		{1, 2, 3, 1},
		{9, 4, 16, 4},
		{16, 30, 6, 6},
	}
	for _, input := range inputs {
		got := Min3(input.a, input.b, input.c)
		if got != input.expected {
			t.Errorf("got: %d; want: %d", got, input.expected)
		}
	}
}
