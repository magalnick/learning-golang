package main

import (
	"flag"
	"fmt"
	"math"
	"strconv"
)

func main() {
	numberToCheck := getNumberToCheck()

	recursive := flag.Bool("recursive", false, "Check all numbers up to the number to check")
	flag.Parse()
	rangeStart := numberToCheck
	if *recursive {
		rangeStart = 2
	}

	for i := rangeStart; i <= numberToCheck; i++ {
		(Number{value: i}).checkIt()
	}
}

/*******************************/
/* Procedural function section */
/*******************************/

func getNumberToCheck(prompts ...string) int {
	if len(prompts) == 0 {
		return getNumberToCheck("Enter a number to check if it's prime: ")
	}
	prompt := prompts[0]
	fmt.Print(prompt)

	var input string
	_, e := fmt.Scanln(&input)
	if e != nil {
		return getNumberToCheck(e.Error())
	}

	numberToCheck, _ := strconv.Atoi(input)
	switch true {

	// check that the string of the original input matches the integer value converted back to a string
	// this protects against both strings and floats being entered
	case strconv.Itoa(numberToCheck) != input:
		return getNumberToCheck("You must enter an integer to check for prime: ")

	case numberToCheck <= 1:
		return getNumberToCheck("Please enter an integer greater than one: ")

	default:
		return numberToCheck
	}
}

/************************/
/* OOP function section */
/************************/

type Number struct {
	value int
	sqrt  int
}

func (n Number) checkIt() {
	if n.isPrime() {
		fmt.Println(n.value, "-> yes")
	} else {
		fmt.Println(n.value, "-> no")
	}
}

func (n Number) isPrime() bool {
	if n.isASimplePrime() {
		return true
	}

	// if up to this value it's not a known simple prime
	// then anything else is specifically not a prime
	// no need to run calculations on any of these values
	if n.value <= 20 {
		return false
	}
	if n.isDivisibleByASimplePrime() {
		return false
	}

	n.calculateSquareRoot()
	for i := 3; i <= n.sqrt; i += 2 {
		if n.value%i == 0 {
			return false
		}
	}

	return true
}

func (n Number) isASimplePrime() bool {
	switch n.value {
	case 2, 3, 5, 7, 11, 13, 17, 19:
		return true
	default:
		return false
	}
}

func (n Number) isDivisibleByASimplePrime() bool {
	v := n.value
	switch true {
	case
		v%2 == 0,
		v%3 == 0,
		v%5 == 0,
		v%7 == 0,
		v%11 == 0,
		v%13 == 0,
		v%17 == 0,
		v%19 == 0:
		return true
	default:
		return false
	}
}

/**
 * Note that this will be an integer of the ceiling value of the square root.
 * If the number is a square with a whole number square root, then it will be exact,
 * however, if the number isn't a square, then the square root will be rounded up
 * to the nearest integer.
 * For example, 24 and 25 will result with 5 as their calculated square root
 * but 26 will result with 6.
 */
func (n *Number) calculateSquareRoot() {
	n.sqrt = int(
		math.Ceil(
			math.Sqrt(
				float64(n.value),
			),
		),
	)
}
