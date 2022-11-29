package main

import "fmt"

func main() {
	fmt.Println("I'm sick of hello world...")

	var whatToSay string
	var i int

	whatToSay = "Goodbye, cruel world"
	fmt.Println(whatToSay)

	i = 7
	fmt.Println("\"i\" is set to", i)

	whatWasSaid, theOtherThingThatWasSaid := saySomething()
	fmt.Printf("The function returned %q and %q\n", whatWasSaid, theOtherThingThatWasSaid)
}

func saySomething() (string, string) {
	return "something", "else"
}
