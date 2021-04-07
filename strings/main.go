package main

import "fmt"

func main() {
	a := "Hello I am here"
	for _, v := range a {
		fmt.Printf("The character %c has value %d\n", v, v)
	}
}
