package exercise

import "testing"

func TestFibo(t *testing.T) {
	expected := []int{
		1, 1, 2, 3, 5, 8, 13, 21,
	}
	for input, want := range expected {
		got := Fibo(input)
		if got != want {
			t.Errorf("got: %d; want: %d", got, want)
		}
	}
}

func TestFact(t *testing.T) {
	expected := []int{
		1, 1, 2, 6, 24, 120,
	}
	for input, want := range expected {
		got := Fact(input)
		if got != want {
			t.Errorf("got: %d; want: %d", got, want)
		}
	}
}

func TestMedian(t *testing.T) {
	type input_struct struct {
		a, b, c, expected int
	}
	inputs := [3]input_struct{
		{1, 2, 3, 2},
		{9, 4, 16, 9},
		{6, 30, 16, 16},
	}
	for _, input := range inputs {
		got := Median(input.a, input.b, input.c)
		if got != input.expected {
			t.Errorf("got: %d; want: %d", got, input.expected)
		}
	}
}
