package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

type Myfloat float64

func (f Myfloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

type Abser interface {
	Abs() float64
}

func main() {
	var a Abser

	v := Vertex{3, 4}
	a = &v
	fmt.Println(a.Abs())

	var b Abser

	b = Myfloat(-math.Sqrt(2))
	fmt.Println(b.Abs())
}
