package main

import "fmt"

func main() {
	cat := "cat"
	if cat == "cat" {
		fmt.Println("cat is cat")
	} else {
		fmt.Println("cat is *not* cat")
	}
}
