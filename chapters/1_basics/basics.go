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
