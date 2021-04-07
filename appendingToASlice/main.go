package main

import "fmt"

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func main() {
	var a []int
	printSlice(a)

	a = append(a, 0)
	printSlice(a)

	a = append(a, 2, 4, 6, 10)
	printSlice(a)

	b := [5]int{1, 2, 3, 4, 10}
	var c []int
	c = b[:]
	c = append(c, 10)
	printSlice(c)
	fmt.Println(b)

}
