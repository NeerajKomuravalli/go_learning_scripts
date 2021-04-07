package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128, 256, 512}

func main() {
	// Multiple ways in which you can use for
	// for i, _ := range pow
	// for _, value := range pow
	// for i := range pow
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}
