// A type assertion provides access to an interface value's underlying concrete value.

package goTour04

import "fmt"

func do(i interface{}) {
	// In a type switch T in i.(T) is replaced with the keyword type.
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

func RunExample15_16() {
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	// Here the type assertion will trigger a panic because the interface value i does not hold the concrete type float64.
	// f = i.(float64) // panic
	// fmt.Println(f)

	// Because we are using ok to validate the assertion, the type assertion will not trigger a panic.
	f, ok := i.(float64)
	fmt.Println(f, ok)

	do(21)
	do("hello")
	do(true)
}
