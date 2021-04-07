package main

import "fmt"

const (
	// Create a huge number by shifting a 1 bit left 100 places.
	// In other words, the binary number that is 1 followed by 100 zeroes.
	Big = 1 << 100
	// Shift it right again 99 places, so we end up with 1<<1, or 2.
	Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }

func needFloat(x float64) float64 {
	return x * 0.1
}
func main() {
	i := 34
	var a float64 = float64(i)
	// Below command will throw error as a and i are of different type and we need explicit conversion
	// var a float64 = i
	t := 3.14
	fmt.Println(i, a, t)
	fmt.Printf("Type : %T and value : %v\n", t, t)
	const Constant = 25
	// You cant do the below, basically you can't reasign a different value to a constant
	// Constant = 100
	// You cannot assign a constant with ":="
	// const constant := 100
	fmt.Printf("Type : %T and value : %v\n", Constant, Constant)
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
	// Below we have converted the Big variale to float64 and that is why it will execute without error as the max value float64 can take is 1.7976931348623157e+308
	fmt.Printf("Big const after being converted to float64 : %v\n", float64(Big))
	// Below command will give "overflowing int" error
	// fmt.Printf("Big const type is : %T\n", Big)
	// Below will throw error becase Big is actually a
	// fmt.Println(needInt(Big))
}
