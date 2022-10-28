package goTour03

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	num0 := 0
	num1 := 0

	return func() int {
		if num0 == 0 {
			if num1 == 0 {
				num1 = 1

				return 0
			}

			num0 = 1
			num1 = 0

			return 1
		} else if num0 == 1 && num1 == 0 {
			num1 = 1
			return 1
		}

		valueToReturn := num0 + num1

		num0 = num1
		num1 = valueToReturn

		return valueToReturn
	}
}

func fibonacciFromGit() func() int {
	f2, f1 := 0, 1
	return func() int {
		// Always return f2 (we return with the next two values already computed)
		// f2 | f1
		//  0 | 1
		//  1 | 1
		//  1 | 2
		//  2 | 3
		//  3 | 5
		//  5 | 8

		// println("OldValues:", f2, f1)

		f := f2
		f2, f1 = f1, f+f1

		// println("NewValues:", f2, f1)
		// println("Solution:", f)

		return f
	}
}

func RunExercise26() {
	f := fibonacciFromGit()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
