package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func main() {
	// The below line is important as without this the m is a nil map (try commenting it and see what happens)
	m := make(map[string]Vertex)
	m["Bell"] = Vertex{20, 40}
	fmt.Println(m["Bell"])

	m2 := make(map[string]string)
	m2["Hello"] = "Bye"
	fmt.Println(m2["Hello"])
}
