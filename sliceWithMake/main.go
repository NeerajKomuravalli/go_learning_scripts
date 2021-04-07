package main

import "fmt"

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}

func main() {
	a := make([]int, 5) // the first parameter is type, second is length
	printSlice("a", a)

	b := make([]int, 0, 5) // // the first parameter is type, second is length, and the third is capacity
	printSlice("b", b)

	c := b[:2]
	printSlice("c", c)

	d := b[2:5]
	printSlice("d", d)
}
