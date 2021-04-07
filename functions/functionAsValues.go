package main

import (
	"fmt"
	"math"
)

func computeFunction(fn func(float64, float64) float64, param1, param2 float64) float64 {
	return fn(param1, param2)
}

func main() {
	calcHypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(calcHypot(5, 12))
	fmt.Println(computeFunction(calcHypot, 3, 4))
	fmt.Println(computeFunction(math.Pow, 3, 4))
}
