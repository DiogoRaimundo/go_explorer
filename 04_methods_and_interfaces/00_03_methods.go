// Go does not have classes. However, you can define methods on types.
// A method is a function with a receiver argument.

// You can only declare a method with a receiver whose type is defined in the same package as the method.
// You cannot declare a method with a receiver whose type is defined in another package (which includes the built-in types such as int).

package goTour04

import (
	"strings"

	"golang.org/x/tour/wc"
)

type MyString string

// (s MyString) is a receiver argument
func (s MyString) Split(separator string) []string {
	return strings.Split(string(s), separator)
}

func WordCountWithMethod(s MyString) map[string]int {
	words := s.Split(" ")
	wordCounter := make(map[string]int)

	for _, word := range words {
		countValue := wordCounter[word]
		wordCounter[word] = countValue + 1
	}

	return wordCounter
}

func WordCount(s string) map[string]int {
	return WordCountWithMethod(MyString(s))
}

func RunExample00_03() {
	wc.Test(WordCount)
}
