package main

import (
	"fmt"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

func (err *MyError) Error() string {
	return fmt.Sprintf("%v Error occured at the time %v", err.What, err.When)
}

func main() {
	merr := &MyError{time.Now(), "Random "}
	merr2 := &MyError{time.Now(), "Random 2 "}
	merr3 := &MyError{time.Now(), "Random 3 "}
	if merr != nil {
		fmt.Println(merr)
	}
	if merr2 != nil {
		fmt.Println(merr2)
	}
	if true {
		fmt.Println(merr3)
	}

}
