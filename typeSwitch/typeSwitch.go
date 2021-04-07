package main

import "fmt"

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("2 * %v = %v\n", v, 2*v)
	case string:
		fmt.Printf("Length of %q is %v\n", v, len(v))
	default:
		fmt.Printf("I don't know about %v\n", v)
	}
}

func main() {
	do(2)
	do("Hello")
	do(22.56)
}
