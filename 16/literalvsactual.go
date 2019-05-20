package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	fileRead := bufio.NewReader(file)
	byteSlice := make([]byte, 0)

	for {
		in, e := fileRead.ReadByte()
		if e == io.EOF {
			break
		}
		if in == byte(' ') || in == byte('\n') {
			continue
		} else {
			byteSlice = append(byteSlice, in)
		}

	}
	fmt.Println(len(byteSlice))

	nextFile, e := os.Open("input.txt")
	if e != nil {
		panic(e)
	}
	newFileReader := bufio.NewReader(nextFile)
	stringSlice := make([]string, 0)
	for {
		inString, err := newFileReader.ReadBytes('\n')
		stringSlice = append(stringSlice, string(inString))
		if err == io.EOF {
			break
		}
	}
	outSlice := make([]string, 0)
	for _, v := range stringSlice {
		v = strings.Replace(v, "\n", "", -1)
		v, e = strconv.Unquote(v)
		outSlice = append(outSlice, v)
	}
	count := 0
	for _, j := range outSlice {
		for range j {
			count++
		}
	}

	anotherFile, _err := os.Open("input.txt")
	if _err != nil {
		panic(err)
	}
	anotherFileReader := bufio.NewReader(anotherFile)

	lineSlice := make([][]byte, 0)
	for i := range lineSlice {
		lineSlice[i] = make([]byte, 0)
	}
	for {
		inString, err := anotherFileReader.ReadBytes('\n')
		inString = bytes.Replace(inString, []byte("\n"), []byte(""), -1)
		lineSlice = append(lineSlice, inString)

		if err == io.EOF {
			break
		}
	}

	var runningCount int
	for _, v := range lineSlice {
		if len(v) == 0 {
			continue
		}
		var w []byte
		for _, y := range v {
			switch {
			case y == byte('"') || y == byte('\\'):
				runningCount += 2
			default:
				w = append(w, y)
				runningCount++
			}
		}
		runningCount += 2
	}
	fmt.Println(runningCount - len(byteSlice))
}
