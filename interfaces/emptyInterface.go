package main

import "fmt"

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func main() {
	var i interface{}
	i = 42
	describe(i)

	i = "Hello"
	describe(i)
}
