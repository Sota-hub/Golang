package main

import (
	"errors"
	"log"
)

func main() {
	result, err := divide(100.0, 0)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("result:", result)
}

func divide(x,y float32) (float32, error) {
	var result float32
	if y == 0 {
		// errors.New(error_message)
		return result, errors.New("Error occurred")
	}
	result = x / y
	return result, nil
}