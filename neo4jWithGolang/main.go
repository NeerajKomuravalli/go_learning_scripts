package main

import (
	"fmt"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

type Data struct {
	Type string
	M    string
	B    string
}

func helloWorld(uri, username, password string) (interface{}, error) {
	driver, err := neo4j.NewDriver(uri, neo4j.BasicAuth(username, password, ""))
	if err != nil {
		return "", err
	}
	defer driver.Close()

	session := driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()

	greeting, err := session.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		data := map[string]interface{}{
			"type": "Select",
			"m":    "Hello",
			"b":    "How are you?",
		}
		data2 := map[string]interface{}{
			"type": "Select",
			"m1":   "I am not fine",
			"b1":   "What about you?",
		}
		result, err := transaction.Run(
			`
			unwind $messages as message
				call apoc.do.when(
					message.type = "Select" and message.m is not null,
					"
					CREATE (a:Greeting) 
						set
							a.m = message.m,
							a.b = message.b
					return a
					",
					"",
					{message:message}
				) yield value as node
			RETURN true
			`,
			map[string]interface{}{"messages": []interface{}{data, data2}})
		if err != nil {
			return nil, err
		}

		if result.Next() {
			return result.Record().Values[0], nil
		}

		return nil, result.Err()
	})
	if err != nil {
		return "", err
	}

	return greeting, nil
}

func main() {
	uri := "bolt://0.0.0.0:7687"
	username := "neo4j"
	password := "orchestral_neo4j"
	greets, err := helloWorld(uri, username, password)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Greets : %v\n", greets)
}
