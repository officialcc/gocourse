package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	FirstName    string  `json:"name"`
	Age          int     `json:"age,omitempty"`
	EmailAddress string  `json:"email,omitempty"`
	Address      Address `json:"address"`
}

type Address struct {
	City  string `json:"city"`
	State string `json:"state"`
}

func main() {

	// person := Person{FirstName: "John", Age: 30, EmailAddress: "sample@example.com"}
	person := Person{FirstName: "John"}

	jsonData, err := json.Marshal(person)
	if err != nil {
		fmt.Println("Error marshalling to JSON:", err)
		return
	}
	fmt.Println(string(jsonData))

	person1 := Person{FirstName: "Jane", Age: 30, EmailAddress: "jane@fakeaddresson.com", Address: Address{City: "New York City", State: "New York"}}

	jasondata1, err := json.Marshal(person1)
	if err != nil {
		fmt.Println("Error marshalling to JSON:", err)
		return
	}
	fmt.Println(string(jasondata1))

	jsonData := {"full_name": "Jane Doe", "emp_id": "0008", "age": 30, "address": {"city": "San Jose", "state": "California"}}
}
