package main

import "log"

type myStruct struct {
	FirstName string
}

// Receiver ties function to type of myStruct
// imagine below
// { FirstName string,
// printFirstName func } 
func (m *myStruct) printFirstName() string {
	return m.FirstName
}

func main() {
	// myVar is like structure which has FirstName
	var myVar myStruct
	myVar.FirstName = "John"

	// Short hand of above
	myVar2 := myStruct{
		FirstName: "Mary",
	}


	log.Println("1:", myVar.printFirstName())
	log.Println("2:", myVar2.printFirstName())
}