package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Hypot() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

type MyFloat64 float64

func (f MyFloat64) Abs() float64 {
	if f < 0 {
		return (float64(-f))
	}
	return float64(f)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println("Hypot result", v.Hypot())

	f := MyFloat64(-math.Sqrt(2))
	fmt.Println("Asolute value is : ", f.Abs())
}
