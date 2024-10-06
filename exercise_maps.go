package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	// return map[string]int{"x": 1}
	str_slice := strings.Fields(s)
	result := make(map[string]int)
	for _, word := range str_slice {
		result[word]++ // by design if key is not exist then map[key] == 0
	}
	return result
}

func main() {
	wc.Test(WordCount)
}
