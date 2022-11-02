package main

import (
	"fmt"
	"strconv"
)

func main() {
	arr := [...]int{10, 20, 30, 40}
	sl1 := arr[1:3]
	sl2 := append(sl1, 50)
	sl3 := append(sl1, 60, 70)
	sl4 := append(sl2, sl3...)
	//sl5 := append(sl4, 1)
	arr2 := [...]int{}
	sl5 := arr2[0:0]
	sl5 = append(sl5, 2, 3)

	fmt.Println(arr, sl1, sl2, sl3, sl4, sl5)

	langCodes := make(map[string]string)
	fmt.Println(len(langCodes))
	fmt.Println(langCodes)
	langCodes["en"] = "English"
	langCodes["he"] = "Hebrew"
	langCodes["de"] = "German"

	langCodes["es"] = "Spanish"
	fmt.Println(langCodes)

	for key, value := range langCodes {
		fmt.Println(key, " => ", value)
	}

	// calling functions
	//callingFunctions()
	pointers()

	fmt.Println("")
}

func pointers() {
	var x, y, z = 1, 2, 3
	var p = &x
	fmt.Printf("x => %v => %v => %v\n", x, &x, *(&x))
	fmt.Printf("y => %v => %v => %v\n", y, &y, *(&y))
	fmt.Printf("z => %v => %v => %v\n", z, &z, *(&z))
	fmt.Printf("p => %v => %v => %v\n", p, &p, *(&p))

	*p = 7
	fmt.Printf("x => %v => %v => %v\n", x, &x, *(&x))
	fmt.Printf("p => %v => %v => %v\n", p, &p, *(&p))

	*p = 1
	fmt.Printf("x => %v => %v => %v\n", x, &x, *(&x))
	fmt.Printf("p => %v => %v => %v\n", p, &p, *(&p))
}

func callingFunctions() {
	integerConcat := func(integers ...int) string {
		var s string
		for _, i := range integers {
			s += strconv.Itoa(i)
		}
		return s
	}
	fmt.Println(integerConcat(0))
}
