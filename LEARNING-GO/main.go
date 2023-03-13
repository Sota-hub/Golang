package main

// Import standard libraries automatically
import "fmt"

// Need at least main func
func main() {
	fmt.Println("Hello, world")

	// Just declare with var
	// Need type of variable after the var name itself
	var whatToSay string
	var i int

	whatToSay = "Goodbye, world"
	fmt.Println(whatToSay);

	i = 7
	fmt.Println("i is set to", i)

	// Short variable declaration 
	// = var whatWasSaid string = saySomething()
	// Only usable inside function
	whatWasSaid, otherSaid := saySomething()
	fmt.Println("The func returned", whatWasSaid, otherSaid)
}

// func name() (return type, return type)
func saySomething() (string, string) {
	return "something", "else"
}