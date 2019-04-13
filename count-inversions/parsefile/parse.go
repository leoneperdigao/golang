package parsefile

import (
	"bufio"
	"io"
	"strconv"
)

//ParseFile function extract integer values insied the text file
func ParseFile(file io.Reader) []int {
	scanner := bufio.NewScanner(file)
	var slice []int
	for scanner.Scan() {
		integer, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic("Something wrong happened. The value is not a integer. Review your file.")
		}
		slice = append(slice, integer)
	}
	if scanner.Err() != nil {
		panic("error happened when reading file")
	}
	return slice
}
