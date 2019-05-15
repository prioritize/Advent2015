package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Range struct {
	X0 int
	X1 int
	Y0 int
	Y1 int
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// Toggle the lights and return modified slice
func toggle(lights [][]bool, locations Range) [][]bool {
	for i := locations.Y0; i <= locations.Y1; i++ {
		innerSlice := lights[i]
		for j := locations.X0; j <= locations.X1; j++ {
			innerSlice[j] = !innerSlice[j]
		}
	}
	return lights
}

// Sum the number of true values and return
func sumBool(lights [][]bool) int {
	fmt.Println(len(lights), len(lights[0]))
	var totalTrue int
	for i := range lights {
		for _, v := range lights[i] {
			if v {
				totalTrue++
			}
		}
	}
	return totalTrue
}

// Turn off
func off(lights [][]bool, locations Range) [][]bool {
	for i := locations.Y0; i <= locations.Y1; i++ {
		innerSlice := lights[i]
		for j := locations.X0; j <= locations.X1; j++ {
			innerSlice[j] = false
		}
	}
	return lights
}

func on(lights [][]bool, locations Range) [][]bool {
	for i := locations.Y0; i <= locations.Y1; i++ {
		innerSlice := lights[i]
		for j := locations.X0; j <= locations.X1; j++ {
			innerSlice[j] = true
		}
	}
	return lights
}

// Construct a Range from a string
func buildRange(start, end string) Range {
	var outRange Range

	startString := strings.Split(start, ",")
	x0, _ := strconv.Atoi(startString[0])
	y0, _ := strconv.Atoi(startString[1])
	endString := strings.Split(end, ",")
	x1, _ := strconv.Atoi(endString[0])
	y1, _ := strconv.Atoi(endString[1])

	outRange.X0 = x0
	outRange.X1 = x1
	outRange.Y0 = y0
	outRange.Y1 = y1
	return outRange
}

func main() {
	file, err := os.Open("input.txt")
	check(err)
	reader := bufio.NewReader(file)

	lights := make([][]bool, 1000)
	for i := range lights {
		lights[i] = make([]bool, 1000)
	}
	for {
		line, e := reader.ReadBytes('\n')
		if e == io.EOF {
			break
		}
		stringLine := strings.Replace(string(line), "\n", "", -1)
		command := strings.Split(stringLine, " ")
		switch {
		case command[0] == "toggle":
			var toggleRange Range
			toggleRange = buildRange(command[1], command[3])
			lights = toggle(lights, toggleRange)
		case command[0] == "turn":
			if command[1] == "off" {
				var offRange Range
				offRange = buildRange(command[2], command[4])
				lights = off(lights, offRange)
			}
			if command[1] == "on" {
				var onRange Range
				onRange = buildRange(command[2], command[4])
				lights = on(lights, onRange)
			}
		default:
			fmt.Println("Something bad happened here")
		}
	}
	fmt.Println(sumBool(lights))

}
