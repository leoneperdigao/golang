package main

import (
	"fmt"
	"os"

	"./countinversions"
	"./parsefile"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		panic("You must pass the file path as parameter. Try again.")
	}

	file, err := os.Open(args[0])
	if err != nil {
		panic(err)
	}
	defer file.Close()

	slice := parsefile.ParseFile(file)

	fmt.Printf("Input size: %d integers.\n", len(slice))
	fmt.Printf("Number of Inversions: %d\n", countinversions.Count(slice))
}
