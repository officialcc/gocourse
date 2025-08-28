package main

import (
	"encoding/xml"
	"fmt"
	"log"
)

type Person struct {
	XMLName xml.Name `xml:"person"`
	Name    string   `xml:"name"`
	Age     int      `xml:"age,omitempty"`
	// City    string   `xml:"city,omitempty"`
	// Email   string   `xml:"email"`
	Email   string   `xml:"-"`
}

func main() {

	// person := Person{Name: "John", Age: 30, City: "London", Email: "email@address.com"}
	person := Person{Name: "John", Email: "email@address.com"}

	xmlData, err := xml.Marshal(person)
	if err != nil {
		log.Fatalln("Error marshalling data into XML:", err) // log.Fatal - no need to return, error will exit program
	}
	fmt.Println(xmlData)

	xmlData1, err := xml.MarshalIndent(person, "", "  ")
	if err != nil {
		log.Fatalln("Error marshalling data into XML:", err)
	}
	fmt.Println(string(xmlData1))

	xmlRaw := `<person><name>Buster</name><age>40</age></person>`

	var personxml Person
	
	err = xml.Unmarshal([]byte(xmlRaw), &personxml)
	if err != nil {
		log.Fatalln("Error unmarshalling XML:", err)
	}
	fmt.Println(personxml)
}
