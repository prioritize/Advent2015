package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

//CheckOthers determines if there are two equal values seperated by one location in a slice
func CheckOthers(line string) bool {
	// Loop through and check line[i] against line[i+2]. If they match return true
	for i := 0; i < len(line)-2; i++ {
		current := line[i]
		skip := line[i+2]
		if current == skip {
			return true
		}
	}
	return false
}

//FindDouble searches a string for additional pairs of values
func FindDouble(line string) bool {
	//Loop through all values in the line.
	for i := 0; i < len(line)-3; i++ {
		// Grab the first two, iterate through all until successful
		pair := line[i : i+2]
		// Set j equal to the next value after the above pair, loop through until the end of line
		for j := i + 2; j < len(line)-1; j++ {
			// Grab the pair indexed by
			value := line[j : j+2]
			// If both pairs match return true
			if pair == value {
				return true
			}
		}
	}
	// Return false if there were not matches found
	return false
}

func main() {
	file, err := os.Open("input.txt")
	check(err)
	reader := bufio.NewReader(file)
	var wordCount int
	for {
		// Read a line
		line, e := reader.ReadBytes('\n')
		if e == io.EOF {
			break
		}
		stringLine := string(line)
		// Remove the trailing newline
		parsedString := strings.Replace(stringLine, "\n", "", -1)
		if CheckOthers(parsedString) {
			if FindDouble(parsedString) {
				wordCount++
			}
		}
	}
	// Report results
	fmt.Println("The total strings that are nice: ", wordCount)
}
