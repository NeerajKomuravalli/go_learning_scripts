package main

import (
	"fmt"
	"strings"
)

func WordCount(s string) map[string]int {
	m := make(map[string]int)
	for _, word := range strings.Fields(s) {
		m[word] += 1
	}
	return m
}

func main() {
	// wc.Test(WordCount)
	fmt.Println("This code will not execute anything but you can copy paste the WordCount function in the go tour to acheive what you want")
}
