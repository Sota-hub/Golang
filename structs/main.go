package main

import "fmt"

type contact struct {
	email string
	zipCode int
}

type person struct {
	firstName string
	lastName string
	contact
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName: "Party",
		contact : contact{
			email: "jim@gmail.com",
			zipCode: 94000,
		},
	}

	jim.updateName("Jimmy")
	jim.print()
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}