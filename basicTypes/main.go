package main

import (
	"fmt"
	"math/cmplx"
)

//variable declaration in blocks
var (
	ToBe bool = false
	//BITSHIFTING In binary, 1, with 64 0s after it (10000...). Same as 2^64.
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
}
