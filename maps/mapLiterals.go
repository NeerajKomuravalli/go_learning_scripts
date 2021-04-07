package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var (
	m = map[string]Vertex{
		"Bell": {2, 10},
		"Hell": Vertex{20, 40}, // This is to show that it works both ways => with and without "Vertex" because we defined it
	}
)

// Above decleration can also be done as shown below
// var m = map[string]Vertex{
// 	"Bell": Vertex{2, 10},
// 	"Hell": Vertex{20, 40},
// }

func main() {
	fmt.Println(m)
}
