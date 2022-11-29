package main

import "fmt"

func main() {
	dog := Dog{
		Name:  "Pidjean",
		Breed: "Not a beagle",
	}

	PrintInfo(&dog)

	gorilla := Gorilla{
		Name:          "Jacque",
		Color:         "Silver",
		NumberOfTeeth: 57,
	}

	PrintInfo(&gorilla)
}

type Animal interface {
	Says() string
	NumberOfLegs() int
}

type Dog struct {
	Name  string
	Breed string
}

func (d *Dog) Says() string {
	return "Yarp"
}

func (d *Dog) NumberOfLegs() int {
	return 4
}

func (g *Gorilla) Says() string {
	return "Grunt"
}

func (g *Gorilla) NumberOfLegs() int {
	return 2
}

type Gorilla struct {
	Name          string
	Color         string
	NumberOfTeeth int
}

func PrintInfo(a Animal) {
	fmt.Println("This animal says", a.Says(), "and has", a.NumberOfLegs(), "legs")
}
