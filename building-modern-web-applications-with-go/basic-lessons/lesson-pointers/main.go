package main

import "log"

func main() {
	var myString string
	myString = "Green"
	log.Printf("myString is set to %q\n", myString)
	changeUsingPointer(&myString)
	log.Printf("After calling the pointer function, myString is set to %q\n", myString)
}

func changeUsingPointer(s *string) {
	log.Println("s is set to", s)
	newValue := "Red"
	*s = newValue
}
