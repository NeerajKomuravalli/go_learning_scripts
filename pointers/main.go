package main

import "fmt"

func main() {
	i, j := 42, 100
	// pointer 'p' is pointing towards 'i'
	p := &i
	// you can also declare p like =>
	// var p *int = &i
	fmt.Printf("Value of i using pointers : %v\n", *p)
	*p = 122
	fmt.Printf("Value of i after changing i through p : %v\n", *p)
	p = &j
	fmt.Printf("Value of j using pointers : %v\n", *p)
	*p = 1100
	fmt.Printf("Value of i after changing j through p : %v\n", *p)
}
