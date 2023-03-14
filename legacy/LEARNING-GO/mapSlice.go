package main

import (
	"log"
	"sort"
)

type User struct {
	First string
	Last string
}

// map is not ordered and not effected by pointer value change
// func main() {
// 	// make map which has string key and int value
// 	myMap := make(map[string]int)
// 	myMap["First"] = 1
// 	myMap["Second"] = 2
// 	log.Println(myMap["First"])
// 	log.Println(myMap["Second"])

// 	// Also able to pass struct
// 	map2 := make(map[string]User)
// 	me := User {
// 		First: "Sota",
// 		Last: "Sasaki",
// 	}
// 	map2["me"] = me
// 	log.Println(map2["me"].First)
// }

// Slice is instead of array in JS
func main() {
	// `[]int` is slice of int 
	var mySlice []int
	mySlice = append(mySlice, 2)
	mySlice = append(mySlice, 1)
	mySlice = append(mySlice, 3)
	// sort is mutable
	sort.Ints(mySlice)
	log.Println(mySlice)
	// output [Sota John]

	// Short hand: []int{items}
	numbers := []int{0,1,2,3,4,5,6,7,8,9}
	log.Println(numbers)
	// 5 to before 8
	log.Println(numbers[5:8])
}
