package fibo

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
