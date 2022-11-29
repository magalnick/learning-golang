package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	myJson := `
[
	{
		"first_name": "Peter",
		"last_name": "Parker",
		"hair_color": "Dirty blond",
		"has_dog": true
	},
	{
		"first_name": "Gwen",
		"last_name": "Stacy",
		"hair_color": "Blond",
		"has_dog": false
	}
]
`

	// to get the values out of the json string,
	// need to "unmarshal" it as array slice into a reference to an interface
	// since it's a reference, it will auto-set the value,
	// any output from the Unmarshal command is an error
	var unmarshalled []Person
	err := json.Unmarshal([]byte(myJson), &unmarshalled)
	if err != nil {
		fmt.Println("Error unmarshalling JSON", err)
	}

	log.Printf("Unmarshalled: %v", unmarshalled)

	// write json from a struct
	var m1 Person
	m1.FirstName = "Mary"
	m1.LastName = "Jane"
	m1.HairColor = "Red"
	m1.HasDog = false

	var m2 Person
	m2.FirstName = "Miles"
	m2.LastName = "Morales"
	m2.HairColor = "Black"
	m2.HasDog = true

	var mySlice []Person
	mySlice = append(mySlice, m1, m2)
	for _, p := range unmarshalled {
		mySlice = append(mySlice, p)
	}

	// Marshal makes JSON
	// MarshallIndent makes pretty
	newJson, err := json.MarshalIndent(mySlice, "", "    ")
	if err != nil {
		fmt.Println("Error marshalling JSON", err)
	}

	// need to convert to a string, otherwise the characters are printed
	// out as their numerical ascii value (eg. space = 32)
	fmt.Println(string(newJson))
}

type Person struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	HairColor string `json:"hair_color"`
	HasDog    bool   `json:"has_dog"`
}
