package main

import (
	"fmt"
)

func Pic(dx, dy int) [][]uint8 {
	im := make([][]uint8, dy)
	for i := range im {
		im[i] = make([]uint8, dx)
	}
	// function using => (x+y)/2
	for y := 0; y < dy; y++ {
		for x := 0; x < dx; x++ {
			im[y][x] = uint8((x + y) / 2)
		}
	}
	return im
}

func main() {
	fmt.Println(Pic(10, 10))
}
