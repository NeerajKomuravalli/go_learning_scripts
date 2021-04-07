package main

import (
	"fmt"
	"sync"
)

type Map struct {
	Mutex sync.Mutex
	M     map[string]int
}

func (m *Map) inc(key string) {
	m.Mutex.Lock()
	m.M[key] += 1
	m.Mutex.Unlock()
}

func main() {
	m := Map{M: make(map[string]int)}
	for i := 0; i <= 100; i++ {
		go m.inc("Hello")
	}
	fmt.Println(m.M["Hello"])
}
