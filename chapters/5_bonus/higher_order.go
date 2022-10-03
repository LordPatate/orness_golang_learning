package main

import (
	"fmt"
	"strings"
)

func mapFunc[T1, T2 any](f func(T1) T2, slice []T1) []T2 {
	newSlice := make([]T2, len(slice))
	for i, elem := range slice {
		newSlice[i] = f(elem)
	}
	return newSlice
}

func main() {
	double := func(x int) int { return x * 2 }
	fmt.Printf("type: %T, value: %v\n", double, double)

	fmt.Println(double(5))

	slice := []int{
		0, 1, 2, 3, 4,
	}
	fmt.Println(mapFunc(double, slice))

	names := []string{
		"Leia", "Luke", "Padme", "Anakin",
	}
	fmt.Println(mapFunc(strings.ToUpper, names))

	var ohMyScope = -1
	fmt.Println(
		mapFunc(
			func(x int) uint { return uint(x + ohMyScope) },
			slice,
		),
	)
}
