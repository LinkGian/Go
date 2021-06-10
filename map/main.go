package main

import "fmt"

func main() {
	//var colors map[int]string

	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
		"white": "#fffff",
	}

	printMap(colors)
}

func printMap(c map[string]string) { // creates "printMap" a map, with i/o of string
	for color, hex := range c { // declares instantiates and creates for loop to iterate through map with "range"
		fmt.Println("Hex code for", color, "is", hex) // will print this text for each value in map above
	}
}
