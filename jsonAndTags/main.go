package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	MiddleName string `json:"middle_name,omitempty"`
	// MiddleName string `json:"middle_name"`
}

func main() {
	jsonString := `{"first_name":"Neeraj", "last_name":"Komuravalli"}`
	// jsonString := `{"first_name":"Ammu", "last_name":"Philip", "middle_name":"Annama"}`
	person := (new(Person))
	json.Unmarshal([]byte(jsonString), person)
	fmt.Println(person)

	newJson, _ := json.Marshal(person)
	fmt.Printf("%s\n", newJson)
}
