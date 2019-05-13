package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// Location is a dim_2 or 2D vector. Used to store locations in santa_map
type Location struct {
	X int
	Y int
}

// NextLocation requires a byte (^, v, <, >) and returns the next location
func NextLocation(move byte, currentLocation Location) Location {
	switch {
	case move == '^':
		//  Up (y + 1)
		currentLocation.Y++

	case move == 'v':
		// Down (y - 1)
		currentLocation.Y--

	case move == '>':
		// Right (x + 1)
		currentLocation.X++

	case move == '<':
		//  Left (x - 1)
		currentLocation.X--
	}
	return currentLocation
}

// CheckExists checks a []Location for existence of a specific Location. If the value exists it returns true, else false
func CheckExists(loc Location, locations []Location) bool {
	for _, v := range locations {
		if (v.X == loc.X) && (v.Y == loc.Y) {
			return true
		}
	}
	return false
}
func check(e error) {
	if e != nil {
		panic(e)
	}
}
func main() {
	var pastLocations []Location
	var teamLocations [2]Location
	var duplicate int
	var turn int

	// var test_location Location
	// Open the input file
	file, err := os.Open("input.txt")
	// Check if there was an error when opening the file
	check(err)
	// Create a NewReader object to read bytes from file
	bufReader := bufio.NewReader(file)
	// The current location
	var currentLocation Location
	for {
		currentRune, e := bufReader.ReadByte()
		if e == io.EOF {

			break
		}
		// If it's Santa's turn to move, access his last location and update based on input
		if turn%2 == 0 {
			teamLocations[0] = NextLocation(currentRune, teamLocations[0])
			currentLocation = teamLocations[0]
		}
		// If it is Robot Santa's turn to move, access his last location and update based on input
		if turn%2 == 1 {
			teamLocations[1] = NextLocation(currentRune, teamLocations[1])
			currentLocation = teamLocations[1]

		}
		// Check if currentLocation exists in pastLocations
		toAppend := CheckExists(currentLocation, pastLocations)
		// If it does not exist, add it to the list
		if !toAppend {
			pastLocations = append(pastLocations, currentLocation)
		}
		// If it does exist, update duplicate (error checking only)
		if toAppend {
			duplicate++
		}
		turn++
	}
	fmt.Println(len(pastLocations))
}
