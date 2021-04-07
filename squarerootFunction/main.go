package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) (float64, int) {
	sqrtOfX := 1.0
	prevSqrtOfX := 0.0
	diff := 100.0
	isNotFinished := true
	iter := 0
	for isNotFinished {
		// The below formula can be derived from Newton method
		sqrtOfX -= (sqrtOfX*sqrtOfX - x) / (2 * sqrtOfX)
		fmt.Println(sqrtOfX)
		if diff-math.Abs(sqrtOfX-prevSqrtOfX) == 0 {
			isNotFinished = false
		} else {
			diff = math.Abs(sqrtOfX - prevSqrtOfX)
			prevSqrtOfX = sqrtOfX
		}
		iter += 1
	}
	return sqrtOfX, iter
}

func main() {
	number := 23104
	sqrtOfX, iter := sqrt(float64(number))
	fmt.Printf("The square rood of %v is %v\nNo of iterations required are : %v\n", number, sqrtOfX, iter)
}
