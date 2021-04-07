package main

import "fmt"

func test() {
	defer fmt.Println("Still Testing!")
	fmt.Println("Testing in the loop!")
}

func main() {
	// This will showcase how does the stacking work with defer with and without the nested function calls
	defer fmt.Print("World!\n")
	fmt.Print("Hello ")
	fmt.Println("Random shit")
	test()
	fmt.Println("\nI will show you how stacking defer works")
	for i := 0; i < 4; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("Now we begine!")
	fmt.Println("You might think it will end here")
}
