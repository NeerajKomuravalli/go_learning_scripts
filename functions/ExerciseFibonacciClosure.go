package main

import "fmt"

// 0, 1, 1, 2, 3, 5, 8, 13, 21, 34
func fibonacci() func() int {
	num1, num2 := 0, 1
	return func() int {
		num := num1
		num1, num2 = num2, num1+num2
		return num
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
