package main

import (
	"log"
	"time"
)

// infer the value type
var s = "seven"

// struct is short for structure
// declare User type
type User struct {
	first string
	last string
	age int
	birthDay time.Time
}

func main() {
	user := User {
		first: "Bruno",
		last: "Suzuki",
	}

	// if user doesn't property value will be default value
	log.Println(user.first, user.last, user.age,user.birthDay)


	var s2 = "six"

	// prioritised than s="seven"
	s := "eight";

	log.Println("s: ",s)
	log.Println("s2: ", s2)

	a,b := saySomething("xxx")
	log.Println("a: ", a, "b: ", b)
}

func saySomething(s3 string) (string, string) {
	return s3, "World"
}

// Lower case is NOT visible outside of package
var something string
// Upper case is visible outside of package
var Something string
// Same as func
func anything() {}
func Anything() {}