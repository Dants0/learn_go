package math

import (
	"math"
)

func Sum(a, b float64) float64 {
	return a + b
}


// Calculate the square root of a number
func Sqrt(x float64) float64 {
	return math.Sqrt(x)
}

// Calculate the factorial of a number
func Factorial(n int) int {
	// if n == 0 {
	// 	return 1
	// }
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	return result
}
