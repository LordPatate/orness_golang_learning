package main

import (
	"fmt"
)

func Hi(name string) string {
	return fmt.Sprintf("Hello, %s!", name)
}

func EuclidianDiv(dividend, divisor int) (quotient, remainder int) {
	remainder = dividend
	for remainder >= divisor {
		remainder -= divisor
		quotient++
	}
	return quotient, remainder
}

func Swap(a, b *int) {
	// a and b are pointers to int

	var tmp int = *a // Verbose declaration and initializing
	/* 	var tmp = *a // Type omitted since it can be infered */
	/* 	tmp := *a // Shorter alternative */
	*a = *b
	*b = tmp
}

func Perimeter(radius float64) float64 {
	const Pi = 3.14 // Untyped constant

	// Pi will take the type float64 as needed by context
	return 2 * Pi * radius
}
