package main

import "fmt"

func main() {
	sum := 1
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
	// in for loop init (i=0) and post statement (i++) are optional.
	// This type of for loop has similar behaviour to while (golang doesn't have while)
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
	// Below command will run forever
	// for {

	// }
}
