package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	data, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("File reading error", err)
		os.Exit(1)
	}
	fmt.Println("Contents of file:", string(data))
}
