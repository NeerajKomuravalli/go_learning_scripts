package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"

	"github.com/google/go-cmp/cmp"
)

type Item struct {
	Title string
	Body  string
}
type API string

var Database []Item

func (api *API) GetItemByName(title string, responseItem *Item) error {
	for _, value := range Database {
		if value.Title == title {
			*responseItem = value
		}
	}
	return nil
}

func (api *API) AddItem(item Item, responseItem *Item) error {
	Database = append(Database, item)
	*responseItem = item
	return nil
}

func (api *API) DeleteItem(item Item, responseItem *Item) error {
	var deleteItem Item
	for idx, value := range Database {
		if cmp.Equal(value, item) {
			Database[idx] = Database[len(Database)-1] //If the order is not important
			Database = Database[:len(Database)-1]
			deleteItem = item
			break
		}
	}
	*responseItem = deleteItem
	return nil
}

func (api *API) EditItem(item Item, responseItem *Item) error {
	var editedItem Item
	for idx, value := range Database {
		if value.Title == item.Title {
			Database[idx].Body = item.Body
			editedItem = item
			break
		}
	}
	*responseItem = editedItem
	return nil
}

func (api *API) GetDB(empty string, responseDB *[]Item) error {
	*responseDB = Database
	return nil
}

func main() {
	api := new(API)
	err := rpc.Register(api)
	rpc.HandleHTTP()
	if err != nil {
		log.Fatal("Error in registering rcp ", err)
	}
	listener, err := net.Listen("tcp", ":4040")
	if err != nil {
		log.Fatal("Listener error", err)
	}
	log.Printf("serving rpc on port %d", 4040)
	err = http.Serve(listener, nil)
	if err != nil {
		log.Fatal("Error serving ", err)
	}
}
