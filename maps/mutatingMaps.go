package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["Hello"] = 1
	fmt.Println(m)

	m["Hello"] = 55
	fmt.Println(m)

	x := m["Hello"]
	fmt.Println(x)

	delete(m, "Hello")
	fmt.Println(m)

	// To check if the value is present or not
	y, ok := m["Bye"]
	fmt.Println(y, ok)
}
