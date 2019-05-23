package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
)

func byteSliceToInt(input []byte) int {
	length := len(input)
	var output int
	iteration := 0
	switch {
	case input[0] == '-':
		for i := length - 1; i > 0; i-- {
			output += int(math.Pow10(iteration)) * int(input[i])
			iteration++
		}
		output = output * -1
	default:
		for i := length - 1; i >= 0; i-- {
			output += int(math.Pow10(iteration)) * int(input[i])
			iteration++
		}
	}
	return output
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	fileReader := bufio.NewReader(file)
	var prior byte
	var total int
	for {
		in, eof := fileReader.ReadByte()
		valid, e := strconv.Atoi(string(in))
		if e == nil {
			var current []byte
			if prior == byte('-') {
				current = append(current, prior)
				current = append(current, byte(valid))
			} else {
				current = append(current, byte(valid))
			}
			for e == nil {
				in, eof = fileReader.ReadByte()
				valid, e = strconv.Atoi(string(in))
				if e == nil {
					current = append(current, byte(valid))
				}
			}
			total += byteSliceToInt(current)
		}
		if eof == io.EOF {
			break
		}
		prior = in

	}
	fmt.Println(total)
}
