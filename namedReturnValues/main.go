package main

import "fmt"

func swap(x, y string) (a, b string) {
	a = y
	b = x
	return
}

func main() {
	fmt.Println(swap("bye!", "Hello,"))
}
