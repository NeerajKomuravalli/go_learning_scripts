package main

import "fmt"

var c, python, java bool
var a, b int = 1, 2
var d, e = "Hi", true

// Below line will throw error. You can initialize a variable like this only inside a function
//f := 100

func main() {
	var i int
	//This is called short variable decleration
	g, h := true, 40
	fmt.Println(i, c, python, java)
	fmt.Println(a, b, d, e)
	fmt.Println(g, h)
}
