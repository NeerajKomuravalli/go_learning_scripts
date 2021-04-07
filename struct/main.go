package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	fmt.Println(Vertex{1, 2})
	v := Vertex{1, 2}
	fmt.Println("Before modification : ", v.X)
	v.X = 22
	fmt.Println("Post modification : ", v.X)

	// Using pointers for doing the same
	p := &v
	fmt.Println("Before modification (using pointers) : ", p.X)
	p.X = 122
	fmt.Println("Post modification (using pointers) : ", p.X)
}
