/*
A function to check whether two binary trees store the same sequence is quite complex in most languages.
*/
package goTour06

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	// Closes the channel when this function returns
	defer close(ch)
	WalkRecursive(t, ch)
}

func WalkRecursive(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}

	// This order of execution is important!
	WalkRecursive(t.Left, ch)
	ch <- t.Value
	WalkRecursive(t.Right, ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go Walk(t1, ch1)
	go Walk(t2, ch2)

	var tValue1, tValue2 int
	ok1, ok2 := true, true
	for ok1 {
		tValue1, ok1 = <-ch1
		tValue2, ok2 = <-ch2

		if ok1 != ok2 || tValue1 != tValue2 {
			return false
		}
	}

	return true
}

func RunExercise07_08() {
	ch := make(chan int)

	go Walk(tree.New(1), ch)

	for v := range ch {
		fmt.Println(v)
	}

	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(1), tree.New(2)))
}
