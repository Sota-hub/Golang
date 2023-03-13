package main

import "fmt"

// methods types
type Animal interface {
	Says() string
	NumberOfLength() int
}

// like a class
type Dog struct {
	Name string
	Breed string 
}

type Gorilla struct {
	Name string
	Color string 
	NumberOfTeeth int
}

func main() {
	dog := Dog{
		Name: "Samson",
		Breed: "German Shephered",
	}

	PrintInfo(&dog)

	gorilla := Gorilla{
		Name: "Jock",
		Color: "grey",
		NumberOfTeeth: 38,
	}

	PrintInfo(&gorilla)
}

func PrintInfo(a Animal) {
	fmt.Println("This animal says", a.Says(), "and has", a.NumberOfLength(), "length")
}

// most receivers are pointer types
func (d *Dog) Says() string {
	return "Woof"
}

func (d *Dog)  NumberOfLength() int {
	return 4
}

func (g *Gorilla) Says() string {
	return "Ugh"
}

func (g *Gorilla)  NumberOfLength() int {
	return 2
}