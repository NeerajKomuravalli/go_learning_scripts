package main

import "fmt"

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}
	// A slice can be declared in either way described below
	var s []int = primes[:4]
	// s := primes[:4]
	fmt.Println(s)
	// As values are passed by referance in slices, any changes made in a slice will affect the underlying array from which the data was read
	s[0] = 25
	fmt.Println("Slice after changes in slice", s)
	fmt.Println("Array after changes in slice", primes)

	// Example of something similar shown above
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names) // [John Paul George Ringo]

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b) // [John Paul] [Paul George]

	b[0] = "XXX"
	fmt.Println(a, b)  // [John XXX] [XXX George]
	fmt.Println(names) // [John XXX George Ringo]

	// Slice literals
	// Below is actually a slice and not an array as the length is not mentioned
	// This first creates an array and then references it to a slice q
	q := []int{1, 2, 3}
	fmt.Println(q)

	bo := []bool{true, false, false}
	fmt.Println(bo)

	st := []struct {
		X int
		Y int
	}{
		{1, 2},
		{4, 5},
	}
	fmt.Println(st)
}
