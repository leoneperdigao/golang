package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
	}
	printMap(colors)
}

func printMap(m map[string]string) {
	for color, hex := range m {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
