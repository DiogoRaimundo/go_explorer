package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	iterCount := 0

	errorDiff := z*z - x
	for math.Abs(errorDiff) > 1e-015 {
		z -= errorDiff / (2 * z)

		errorDiff = z*z - x
		iterCount++
	}

	fmt.Printf("Found solution in %d itearions with error %e\n", iterCount, errorDiff)
	return z
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(math.Sqrt(2))
}
