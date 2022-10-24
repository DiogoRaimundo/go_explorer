package main

import (
	"fmt"
	"math"
)

// Go programs express error state with error values. The error type is a built-in interface similar to fmt.Stringer.
// Functions often return an error value, and calling code should handle errors by testing whether the error equals nil.
type error interface {
	Error() string
}

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}

	z := 1.0
	iterCount := 0

	errorDiff := z*z - x
	for math.Abs(errorDiff) > 1e-015 {
		z -= errorDiff / (2 * z)

		errorDiff = z*z - x
		iterCount++
	}

	fmt.Printf("Found solution in %d itearions with error %e\n", iterCount, errorDiff)
	return z, nil
}

func main() {
	// Error example
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
