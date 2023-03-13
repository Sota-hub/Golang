package main

import (
	"log"
	"github.com/Sota-hub/Golang/LEARNING-GO/packages/helpers"
)

// func main() {
// 	// import original package and use it
// 	var myVar helpers.SomeType
// 	myVar.TypeName = "Some name"
// 	log.Println(myVar.TypeName)
// }

// intChan is type of chanel of int
func CalculateValue(intChan chan int) {
	randomNumber := helpers.RandomNumber(100)
	// send number into channel (like await)
	intChan <- randomNumber
}

func main() {
	// channel is used to communicate multi goroutines
	// create channel 
	intChan := make(chan int)

	// defer execute as soon as current func(=main()) is done = close channel after main() is executed
	// write 
	defer close(intChan)

	// go launch new goroutine which move like async
	go CalculateValue(intChan)

	// assign number from channel (like await)
	num := <- intChan
	log.Println(num)
}