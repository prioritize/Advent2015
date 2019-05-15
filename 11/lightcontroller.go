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
func toggle(lights [][]bool, locations Range) [][]bool {
	for i := locations.Y0; i < locations.Y1; i++ {
		for j := locations.X0; j < locations.X1; j++ {
			lights[i][j] = !lights[i][j]
		}
	}
	return lights
}

func sumBool(lights [][]bool) int {
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

func off(lights [][]bool, locations Range) [][]bool {
	for i := locations.Y0; i < locations.Y1; i++ {
		for j := locations.X0; j < locations.X1; j++ {
			lights[i][j] = false
		}
	}
	return lights
}

func on(lights [][]bool, locations Range) [][]bool {
	for i := locations.Y0; i < locations.Y1; i++ {
		for j := locations.X0; j < locations.X1; j++ {
			lights[i][j] = true
		}
	}
	return lights
}
func buildRange(x0, x1, y0, y1 int) Range {
	var outRange Range
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
			// TODO: Place this into a function that gets passed the two strings prior to parsing and  returns the ranget st
			startString := strings.Split(command[1], ",")
			x0, _ := strconv.Atoi(startString[0])
			y0, _ := strconv.Atoi(startString[1])
			endString := strings.Split(command[3], ",")
			x1, _ := strconv.Atoi(endString[0])
			y1, _ := strconv.Atoi(endString[1])
			toggleRange = buildRange(x0, x1, y0, y1)
			lights = toggle(lights, toggleRange)
		}
	}
	fmt.Println(sumBool(lights))

}
