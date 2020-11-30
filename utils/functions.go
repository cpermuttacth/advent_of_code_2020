package utils

import (
	"fmt"
)

// Greatest Common Denominator
func Gcd(a, b int64) int64 {
	for b != 0 {
		a, b = b, a%b
	}

	return a
}

// Least Common Multiple
func Lcm(a, b int64) int64 {
	return a * b / Gcd(a, b)
}

// Kernighan's Bit Counting Algorithm
func CountBits(n uint32) int {
	count := 0
	for n > 0 {
		n = n & (n - 1)
		count++
	}

	return count
}

// Check if error is not nil and panic with message if it is.
func Check(e error, message string) {
	if e != nil {
		panic(fmt.Errorf("%s: %s", message, e))
	}
}