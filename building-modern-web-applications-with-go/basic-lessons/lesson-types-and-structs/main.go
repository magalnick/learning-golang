package main

import (
	"fmt"
	"time"
)

func main() {
	user := User{
		FirstName:   "Bob",
		LastName:    "Bobberson",
		PhoneNumber: "(858) 858-8558",
		Age:         57,
		BirthDate:   time.Time{},
	}

	user2 := User{FirstName: "Tim"}

	fmt.Println(user.getFirstName())
	fmt.Println(user2.getFirstName())
}

type User struct {
	// Note that these are capitalized to make them public.
	// If they are lowercase, they're private to only this package.
	// Interesting...
	FirstName   string
	LastName    string
	PhoneNumber string
	Age         int
	BirthDate   time.Time
}

func (u *User) getFirstName() string {
	return u.FirstName
}
