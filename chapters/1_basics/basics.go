package main

import (
	"fmt"
)

func Hi(name string) string {
	return fmt.Sprintf("Hello, %s!", name)
}

func EuclidianDiv(dividend, divisor int) (quotient, remainder int) {
	for dividend >= (quotient+1)*divisor {
		quotient++
	}
	remainder = dividend - quotient*divisor
	return quotient, remainder
}
