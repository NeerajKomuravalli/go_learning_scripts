package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Microsecond)
		fmt.Println(s)
	}
}

func sayMore(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(50 * time.Microsecond)
		fmt.Println(s)
	}
}

func main() {
	// Code will run the "go say("World!")" statement and then move to the next one and when the execution is done it will end the programme
	go say("World!")
	sayMore("Hello")

}
