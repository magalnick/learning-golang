package helpers

import (
	"math/rand"
	"time"
)

func RandomNumber(n int) int {
	// this random function gives back an integer from 0 to n-1
	// so if n = 10, the random numbers will be 0 - 9
	// however if the rand isn't pre-seeded with something to force it to actually be more random,
	// the same value will always result every time the program runs
	rand.Seed(time.Now().UnixNano())
	value := rand.Intn(n * n)
	return value
}
