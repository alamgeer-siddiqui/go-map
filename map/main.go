package main

import (
	"fmt"
)

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "dsd099",
		"white": "#fffff",
	}

	printMap(colors)

	// var colors map[string]string

	// colors:=make(map[int]string)
	// colors[10]="#ff4443"
	// delete(colors,10)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("hex code for ", color, "is ", hex)
	}
}
