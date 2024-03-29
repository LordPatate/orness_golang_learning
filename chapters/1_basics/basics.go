package basics

import (
	"fmt"
)

func Hi(name string) string {
	return fmt.Sprintf("Hello, %s!", name)
}

func EuclidianDiv(dividend, divisor int) (quotient, remainder int) {
	// quotient and remainder are named return values
	remainder = dividend
	for remainder >= divisor {
		remainder -= divisor
		quotient++
	}
	return // Same as `return quotient, remainder`
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

func Min3(a, b, c int) int {
	if b > c {
		b = c
	}
	if a > b {
		a = b
	}
	return a
}

func Countdown(n int) {
	for i := n; i > 0; i-- {
		fmt.Println(i)
	}
	// 'i' is no longer in scope
	fmt.Println("🚀 To inifinity, and beyond!")
}
