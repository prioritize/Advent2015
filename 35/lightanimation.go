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
	disp := display.NewDisplay(display.NewPixel(100, 100))
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
	disp.ReplaceDisplay(lights)
	// count := 0
	// for i := 0; i < 100; i++ {
	// 	for j := 0; j < 100; j++ {
	// 		fmt.Println(disp.Neighbors(display.NewPixel(i, j)))
	// 		count++
	// 	}
	// }
	// fmt.Println(count)
	i := 0
	for i < 100 {
		disp.Animate()
		i++
	}
	fmt.Println(disp.SumDisplay())
	// fmt.Println(disp.GetDisplay())

	// fmt.Println(lights)
	// fmt.Println(len(lights))
	// fmt.Println(len(lights[0]))
}
