package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	fmt.Println("Go lang runs on?")
	// switch with a constant
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X")
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Println(os)
	}

	fmt.Println("When is Saturday?")
	today := time.Now().Weekday()
	// switch with a variable whose type can be anything
	switch time.Saturday {
	case today + 0:
		fmt.Println("It's today")
	case today + 1:
		fmt.Println("It's tomorrow")
	case today + 2:
		fmt.Println("It's day after")
	default:
		fmt.Println("Too far!")
	}

	fmt.Print("Please wish uncle ")
	t := time.Now()
	// Switch with no condition => works like switch true => this will implement the first true case it will find and ignore the rest
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon!")
	default:
		fmt.Println("Good evening!")
	}
}
