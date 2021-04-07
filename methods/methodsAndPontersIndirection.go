package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Scale(f float64) {
	v.X *= f
	v.Y *= f
}

func (v Vertex) Scale2(f float64) {
	v.X *= f
	v.Y *= f
}

func DiffScale(v *Vertex, f float64) {
	v.X *= f
	v.Y *= f
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func AbsFunc(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	v.Scale(2)
	DiffScale(&v, 10)

	p := &Vertex{4, 3}
	p.Scale(3)
	DiffScale(p, 8)

	fmt.Println(v, *p)

	v1 := Vertex{3, 4}
	fmt.Println(v1.Abs())
	fmt.Println(AbsFunc(v1))

	p1 := &Vertex{4, 3}
	fmt.Println(p1.Abs())
	fmt.Println(AbsFunc(*p1))
}
