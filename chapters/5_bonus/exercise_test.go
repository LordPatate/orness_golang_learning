package exercise

import "testing"

func TestFold(t *testing.T) {
	slice := []int{
		1, 2, 3, 4,
	}
	sum := Fold(slice, func(a, b int) int { return a + b }, 0)
	prod := Fold(slice, func(a, b int) int { return a * b }, 1)
	if sum != 10 {
		t.Errorf("got %v, want %v\n", sum, 10)
	}
	if prod != 24 {
		t.Errorf("got %v, want %v\n", prod, 24)
	}
}
