package main

import "fmt"

func Fact(n int) (int, error) {
	if n < 0 {
		return 0, fmt.Errorf("n must be positive but is %d", n)
	}
	if n < 2 {
		return 1, nil
	}
	x, err := Fact(n - 1)
	if err != nil {
		return 0, err
	}
	return n * x, nil
}
