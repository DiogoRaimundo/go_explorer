// Go also supports methods with pointer receivers.
// A method is a function with a receiver argument.

package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

// Function with a value receiver
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Function with a pointer receiver
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// Function with a value receiver
// Since AbsIfScale has a value receiver, if operates on a copy of the original Vertex value.
// This is the same behavior as for any other function argument.
func (v Vertex) AbsIfScale(f float64) float64 {
	v.X = v.X * f
	v.Y = v.Y * f

	return v.Abs()
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())

	// Note that this function return the absolute value of scaled Vertex without scaling the passed Vertex instance
	fmt.Println(v.AbsIfScale(10))

	v.Scale(10)
	fmt.Println(v.Abs())

	/*
		Also note that both Abs and Scale are called using v, even though Scale receives a pointer and v is a value.

		While functions with a pointer argument must take a pointer, methods with pointer receivers can take either a value or a pointer
		This is because, as convenience, Go interprets the statement v.Scale(5) as (&v).Scale(5).

		This behaviour is also apply for value receivers.
		While functions that take a value argument must take a value of that specific type, methods with value receivers can take either a value or a pointer.
		Go interprets the statement v.Abs() as (*v).Abs().
	*/

	/*
		When choosing between a value or pointer receiver, take in consideration that a pointer receivers:
		- Allows the method to modify the value that its receiver points to;
		- Avoids copying the value on each method call. This can be more efficient if the receiver is a large struct, for example.

		In general, all methods on a given type should have either value or pointer receivers, but not a mixture of both.
	*/
}
