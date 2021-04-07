package main

import "fmt"

func main() {
	c := make(chan int, 2)
	c <- 2
	c <- 10
	// Doing the below will throw error as the buffee length is just 2
	// c <- 22
	fmt.Println(<-c)
	fmt.Println(<-c)
	// Doing the below will throw error as the buffee length is just 2
	// fmt.Println(<-c)
}
