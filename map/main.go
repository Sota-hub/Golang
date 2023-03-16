package main

import "fmt"

type colorMap map[string]string

func main() {
	// var colors map[string]string

	// colors := make(map[string]string)
	// colors["white"] = "#ffffff"
	// delete(colors, "white")

	colors := colorMap {
		"red": "#ff0000",
		"green": "#4b745",
		"white": "#ffffff",
	}

	printMap(colors)
}

func printMap(c colorMap) {
	for key,value := range c {
		fmt.Println("Hex code for", key, "is", value)
	}
}