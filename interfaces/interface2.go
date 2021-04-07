package main

import "fmt"

type T struct {
	s string
}

func (t T) M() {
	fmt.Println(t.s)
}

type I interface {
	M()
}

func main() {
	t := T{"Hello"}
	t.M()
}
