package goTour02

import (
	"fmt"
	"math"
)

func sqrt(x float64) float64 {
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

func RunExample08() {
	fmt.Println(sqrt(2))
	fmt.Println(math.Sqrt(2))
}
