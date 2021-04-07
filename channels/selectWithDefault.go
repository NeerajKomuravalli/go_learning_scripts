package main

import (
	"fmt"
	"time"
)

func main() {
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)

	for {
		select {
		case <-tick:
			fmt.Print("Tick... ")
		case <-boom:
			fmt.Print("Boom!\n")
			return
		default:
			time.Sleep(500 * time.Microsecond)
		}
	}
}
