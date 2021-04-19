package main

import (
	"fmt"
	"log"
	"net/rpc"
)

type Item struct {
	Title string
	Body  string
}

func main() {
	var item Item
	var database []Item

	client, err := rpc.DialHTTP("tcp", "localhost:4040")
	if err != nil {
		log.Fatal("Connection error ", err)
	}

	client.Call("API.AddItem", Item{"1", "One"}, &item)
	client.Call("API.AddItem", Item{"2", "Two"}, &item)
	client.Call("API.AddItem", Item{"3", "Three"}, &item)
	client.Call("API.EditItem", Item{"2", "Two (Edited)"}, &item)
	client.Call("API.DeleteItem", Item{"1", "One"}, &item)
	client.Call("API.GetDB", "", &database)
	fmt.Println(database)
}
