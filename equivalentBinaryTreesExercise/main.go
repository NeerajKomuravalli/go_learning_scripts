package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	defer close(ch)
	var walker func(t *tree.Tree, ch chan int)
	walker = func(t *tree.Tree, ch chan int) {
		if t == nil {
			return
		}
		walker(t.Left, ch)
		ch <- t.Value
		walker(t.Right, ch)
	}
	walker(t, ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	for {
		v1, ok1 := <-ch1
		v2, ok2 := <-ch2
		if ok1 == false || ok2 == false {
			break
		}
		if v1 != v2 {
			return false
		}
	}
	return true
}

func main() {
	t := tree.New(1)
	ch := make(chan int)
	go Walk(t, ch)
	for v := range ch {
		fmt.Println(v)
	}
	t1 := tree.New(1)
	t2 := tree.New(10)
	fmt.Println(Same(t1, t2))
}
