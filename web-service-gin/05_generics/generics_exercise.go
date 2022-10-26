/*
	Go functions can be written to work on multiple types using type parameters.
	The type parameters of a function appear between brackets, before the function's arguments.
	Go also supports generic types by declaring it the same way.
	A type can be parameterized with a type parameter, which could be useful for implementing generic data structures.
*/

package main

import "fmt"

// Index returns the index of x in s, or -1 if not found.
// comparable is a useful constraint that makes it possible to use the == and != operators on values of the type.
func Index[T comparable](s []T, x T) int {
	for i, v := range s {
		// v and x are type T, which has the comparable
		// constraint, so we can use == here.
		if v == x {
			return i
		}
	}
	return -1
}

// List represents a singly-linked list that holds
// values of any type.
type List[T any] struct {
	next *List[T]
	val  T
}

func (l *List[T]) AddNext(nextItem *List[T]) {
	for l.next != nil {
		l = l.next
	}

	l.next = nextItem
}

func main() {
	// Index works on a slice of ints
	si := []int{10, 20, 15, -10}
	fmt.Println(Index(si, 15))

	// Index also works on a slice of strings
	ss := []string{"foo", "bar", "baz"}
	fmt.Println(Index(ss, "hello"))

	firstItem := &List[int]{val: 2}
	secondItem := &List[int]{val: 4}
	thirdItem := &List[int]{val: 6}
	fourthItem := &List[int]{val: 8}

	firstItem.AddNext(secondItem)
	firstItem.AddNext(thirdItem)
	firstItem.AddNext(fourthItem)

	item := firstItem
	fmt.Println(item.val)
	for item.next != nil {
		item = item.next
		fmt.Println(item.val)
	}
}
