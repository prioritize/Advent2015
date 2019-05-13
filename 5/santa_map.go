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
		fmt.Println("Case up")
		currentLocation.Y++
		//  Up (y + 1)

	case move == 'v':
		// Down (y - 1)
		fmt.Println("Case down")
		currentLocation.Y--

	case move == '>':
		// Right (x + 1)
		fmt.Println("Case right")
		currentLocation.X++

	case move == '<':
		//  Left (x - 1)
		fmt.Println("Case left")
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
	var locationSlice []Location
	var duplicate int
	var reads int
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
		currentLocation = NextLocation(currentRune, currentLocation)
		toAppend := CheckExists(currentLocation, locationSlice)
		if !toAppend {
			locationSlice = append(locationSlice, currentLocation)
			reads++
		}
		if toAppend {
			duplicate++
		}
		fmt.Println(currentLocation)

	}
	fmt.Println(currentLocation)
	// fmt.Println(locationSlice)
	fmt.Println(duplicate)
	fmt.Println(len(locationSlice))
	fmt.Println(reads)

}
