package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {
	result, err := divide(100, 0)
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println("Result of division is", result)
}

func divide(x, y float64) (float64, error) {
	var result float64

	if y == 0 {
		return result, errors.New("Cannot divide by 0")
	}
	result = x / y
	return result, nil
}
