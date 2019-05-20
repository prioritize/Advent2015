package main

import (
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

	fileRead := bufio.NewReader(file)
	byteSlice := make([]byte, 0)

	for {
		in, e := fileRead.ReadByte()
		if e == io.EOF {
			break
		}
		byteSlice = append(byteSlice, in)
	}
	fmt.Println(len(byteSlice))

	fileRead.Reset()

}
