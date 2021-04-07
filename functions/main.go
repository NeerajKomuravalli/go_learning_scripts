package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func add2(x, y int) int {
	return x + y
}

func getMultipleResults(x, y string) (string, string) {
	return y, x
}

func main() {
	fmt.Println(add(10, 100))
	fmt.Println(add2(20, 40))
	fmt.Println(getMultipleResults("hello", "bye"))
}
