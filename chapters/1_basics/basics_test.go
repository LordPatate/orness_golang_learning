package main

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
			t.Errorf("got:%d, %d; want: %d, %d", q, r, input.expected_q, input.expected_r)
		}
	}
}
