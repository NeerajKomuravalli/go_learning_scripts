package main

import "fmt"

type Vertex struct {
	X, Y int
}

var (
	v1 = Vertex{1, 2}
	// Y = 0
	v2 = Vertex{X: 22}
	// X = 0
	v3 = Vertex{Y: 44}
	// Both X and Y are 0
	v4 = Vertex{}
	// Pointing towards the memeory holding {1, 2}
	p = &Vertex{1, 2}
)

func main() {
	v1.X = 44
	v1.Y = 22
	p.X = 111
	p.Y = 144
	fmt.Println(v1, v2, v3, v4, p, *p)
}
