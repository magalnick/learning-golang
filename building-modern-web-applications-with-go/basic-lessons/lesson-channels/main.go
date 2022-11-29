package main

/**
 * This program uses channels to run things concurrently.
 * The "defer" is also used here to close the channel, since it's an open pointer.
 * The "defer" call says do this thing at the end of the current function call,
 * so it will be used for things like closing database connections or file pointers
 * immediately after opening them so that you don't forget to do it.
 *
 * Putting "go" before a command is what has it run concurrently (asynchronously).
 * Then to listen for the channel value, use the left arrow.
 * Put a value into the channel with the left arrow, then get the value out later on with the left arrow.
 */

import (
	"github.com/magalnick/learning-golang/building-modern-web-applications-with-go/lesson-channels/helpers"
	"log"
)

const numPool = 10

func main() {
	intChan := make(chan int)
	defer close(intChan)

	for i := 0; i < 1; i++ {
		go CalculateValue(intChan)
		num := <-intChan
		log.Println(num)
	}
}

func CalculateValue(intChan chan int) {
	randomNumber := helpers.RandomNumber(numPool)
	intChan <- randomNumber
}
