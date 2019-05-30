package main

import (
	display "Advent2015/35/lightmatrix"
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	fileReader := bufio.NewReader(file)
	lights := make([][]int, 0)
	for i := range lights {
		lights[i] = make([]int, 0)
	}
	lightRow := make([]int, 0)
	for {
		l, e := fileReader.ReadByte()
		if e == io.EOF {
			break
		}
		switch {
		case string(l) == "#":
			lightRow = append(lightRow, 1)
		case string(l) == ".":
			lightRow = append(lightRow, 0)
		case string(l) == "\n":
			lights = append(lights, lightRow)
			lightRow = make([]int, 0)
		default:
			fmt.Println("Default")
		}
	}
	disp := display.NewDisplay(display.NewPixel(len(lights[0]), len(lights)))
	disp.ReplaceDisplay(lights)
	iterations := 0
	for iterations < 100 {
		disp.Animate(false)
		iterations++
	}
	fmt.Println(disp.SumDisplay())
}
