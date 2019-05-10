package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func wrapping_paper_area(length, width, height int) {
	// side_1 := length * width
	// side_2 := length * height
	// side_3 := height * width

}
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	var area int
	f, err := os.Open("input.txt")
	check(err)
	buf_file := bufio.NewReader(f)
	for {
		buffered_line, err := buf_file.ReadBytes('\n')
		if err == io.EOF {
			break
		}
		present := string(buffered_line)
		dim_3 := strings.Split(present, "x")

		fmt.Println(dim_3)
		fmt.Println(present)
	}

	fmt.Println("The area required for all these damn presets is", area)

}
