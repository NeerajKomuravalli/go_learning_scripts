package main

import "fmt"

func main() {
	// Declaration type 1
	var a [2]int
	a[0] = 22
	fmt.Println(a)
	a[1] = 55
	fmt.Println(a)
	fmt.Println(a[0], a[1])

	// Declaration type 2
	p := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(p)
}
