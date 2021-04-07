package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

// if with short statement (v:=math.Pow(x, n))
func powWithLimit(x, n, limit float64) float64 {
	if v := math.Pow(x, n); v < limit {
		return v
	} else {
		fmt.Printf("Value %v > %v (limit)\n", v, limit)
	}
	return limit
}

func main() {
	fmt.Printf("%v \n%v\n", sqrt(2), sqrt(-4))
	fmt.Println(powWithLimit(2, 9, 1024))
}
