package main

import "log"

func main() {
	var myString string
	myString = "Green"

	log.Println("before func:", myString)

	// `&` points variable memory location
	changeUsingPointer(&myString);

	log.Println("after func:", myString)
}

// `*` means pointer
// type is pointer pointed value is string
func changeUsingPointer(s *string) {
	log.Println("memory location:", s)
	newValue := "Red"
	// rewrite value of pointed
	*s = newValue
}