package main

import (
	"fmt"
)

func main() {
	// var colors map[string]string

	// colors := make(map[string]string)

	colors := map[string]string{
		"red": "#ff0000",
		"green": "#ff3333",
	}

	colors["white"]  = "#ffffff"

	printMap(colors)
	// fmt.Println(colors)
}

func printMap (c map[string]string) {
	for col, hex := range c {
		fmt.Println("hex code for", col, "is", hex)
	}
}