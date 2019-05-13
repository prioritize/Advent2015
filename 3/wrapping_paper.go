package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"strconv"
	"sort"
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
func process_line(line []byte)([]int){
	// process_line accepts a []byte (typically from ReadBytes) and returns a sorted []int
	var out_int []int 
	// Cast the buffered_line that is a rune into a string
	string_line := string(line)
	// Remove the '\n' and replace it with an empty character. The ReadBytes() command includes the delimiter rune
	remove_newline := strings.Replace(string_line, "\n", "", -1)
	//Split the string into a slice of three values
	split_x:= strings.Split(remove_newline, "x")
	// For each of the items in split_x, convert from string to int, and append to an []int
	for _, v := range split_x{
		int_v, err := strconv.Atoi(v)
		// Check for errors and close program if an error is encountered
		if err != nil{
			fmt.Println(err)
			os.Exit(2)
		}
		out_int = append(out_int, int(int_v))
	}
	// Sort the []int which will enables area calculation
	sort.Ints(out_int)
	return out_int	
}
func paper_area_required(box_dims []int)(int){
	return 3*(box_dims[0] * box_dims[1]) + 2 * (box_dims[1] * box_dims[2]) + 2 * (box_dims[0] * box_dims[2])

}

func main() {
	// The area to calculate and return
	var area int = 0
	// Open the input file
	f, err := os.Open("input.txt")
	// Check if t here was an error
	check(err)
	// Feed the opened file into a NewReader object so that we can use ReadBytes()
	buf_file := bufio.NewReader(f)
	length := 0
	for {
		// Read a line until a '\n' is found
		buffered_line, err := buf_file.ReadBytes('\n')
		// Check if there was an error while using the ReadBytes command
		if err == io.EOF {
			break
		}
		box_dims := process_line(buffered_line)
		area += paper_area_required(box_dims)
		length += 1
	}
	fmt.Println("The area required for all these damn presets is", area)

}
