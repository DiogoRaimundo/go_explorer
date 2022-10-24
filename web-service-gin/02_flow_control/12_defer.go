// A defer statement defers the execution of a function until the surrounding function returns.
// The deferred call's arguments are evaluated immediately, but the function call is not executed until the surrounding function returns.
// Deferred function calls are pushed onto a stack. When a function returns, its deferred calls are executed in last-in-first-out order.

// More information about defers (https://go.dev/blog/defer-panic-and-recover)

package main

import "fmt"

func main() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}
