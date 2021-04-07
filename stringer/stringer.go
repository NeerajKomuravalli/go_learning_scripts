package main

import "fmt"

type personalDetails struct {
	Name string
	Id   int
}

func (d personalDetails) String() string {
	return fmt.Sprintf("My name is %v and my id is %v\n", d.Name, d.Id)
}

func main() {
	data := personalDetails{"Neeraj", 2345}
	// Can also be declared like below
	// data := personalDetails{
	// 	Name: "Neeraj",
	// 	Id:   2345,
	// }
	fmt.Println(data)
}
