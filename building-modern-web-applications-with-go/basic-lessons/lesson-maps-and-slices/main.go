package main

import (
	"fmt"
	"sort"
)

func main() {
	myMap := make(map[string]User)

	bob := User{
		FirstName: "Bob",
		LastName:  "Bobberson",
	}
	myMap["bob"] = bob

	fmt.Println(myMap)

	var names []string
	names = append(names, "Jim")
	names = append(names, "Bob")
	names = append(names, "Tim")
	names = append(names, "Sally")
	fmt.Println(names)
	sort.Strings(names)
	fmt.Println(names)
	fmt.Println(names[1:3])
}

type User struct {
	FirstName string
	LastName  string
}
