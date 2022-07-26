package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	frequency := map[string]int{}
	for _, character := range strings.Split(s, " ") {
		frequency[character]++
	}

	return frequency
}

func main() {
	wc.Test(WordCount)
}
