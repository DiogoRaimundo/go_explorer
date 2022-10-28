// An interface type is defined as a set of method signatures and can hold any value that implements those methods.
// A type implements an interface by implementing its methods (without explicit declaration of intent).
// This decouples the definition of an interface from its implementation.

package goTour04

import (
	"fmt"
	"math"
)

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

type F float64

func (f F) M() {
	fmt.Println(f)
}

func RunExample09_14() {
	var i I

	i = &T{"Hello"}
	describe(i)
	i.M()

	i = F(math.Pi)
	describe(i)
	i.M()

	// Here, i is a T (not *T) and does NOT implement I.
	// i = T{"Hello"}
	// describe(i)
	// i.M()

	// If the concrete value inside the interface itself is nil, the method will be called with a nil receiver.
	// This allows writing methods that gracefully handle being called with a nil receiver.
	var t *T
	i = t
	describe(i)
	i.M()

	// A nil interface value holds neither value nor concrete type.
	// Calling a method on a nil interface is a run-time error.
	// There is no type inside the interface tuple to indicate which concrete method to call.
	var j I
	describe(j)
	// j.M()
}

// The interface type that specifies zero methods is known as the empty interface which may hold values of any type.
// Empty interfaces are used by code that handles values of unknown type.
func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
