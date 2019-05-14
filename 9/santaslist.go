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

//CheckDoubles steps through the string and checks for doubles
func CheckDoubles(line string) bool {
	for i := 0; i < len(line)-1; i++ {
		current := line[i]
		next := line[i+1]
		if current == next {
			// If a pair is found return true
			return true
		}
	}
	// No pairs found, return false
	return false
}

// CheckSpecifics checks the specific values listed in Day 5
func CheckSpecifics(line string) bool {
	// The specific values listed
	first := "ab"
	second := "cd"
	third := "pq"
	fourth := "xy"
	// Check if the specific values exist. If exists return false
	for i := 0; i < len(line)-1; i++ {
		pair := line[i : i+2]
		switch {
		case pair == first:
			return false
		case pair == second:
			return false
		case pair == third:
			return false
		case pair == fourth:
			return false
		}
	}
	// If none of the specifics exist return true
	return true
}

// CountVowels counts the total number of vowels in the string supplied
func CountVowels(line string) bool {
	count := 0
	for _, v := range line {
		switch {
		case v == 'a':
			count++
		case v == 'e':
			count++
		case v == 'i':
			count++
		case v == 'o':
			count++
		case v == 'u':
			count++
		}
	}
	// If 3 or more vowels exist this string passes
	if count > 2 {
		return true
	}
	// Less than 3 vowels fails
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
		if CheckSpecifics(parsedString) {
			if CountVowels(parsedString) {
				if CheckDoubles(parsedString) {
					wordCount++
				}
			}
		}
	}
	// Report results
	fmt.Println("The total strings that are nice: ", wordCount)
}
