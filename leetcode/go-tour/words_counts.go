package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	words := strings.Fields(s)
	counts := make(map[string]int)

	for i := 0; i < len(words); i++ {
		counts[words[i]] += 1
	}

	return counts
}

func main() {
	wc.Test(WordCount)
}
