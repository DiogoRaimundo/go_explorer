package goTour03

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	words := strings.Split(s, " ")
	wordCounter := make(map[string]int)

	for _, word := range words {
		countValue := wordCounter[word]
		wordCounter[word] = countValue + 1
	}

	return wordCounter
}

func RunExercise23() {
	wc.Test(WordCount)
}
