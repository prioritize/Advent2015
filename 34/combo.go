package main

import (
	"Advent2015/33/combinations"
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	// comb := combinations.New([]string{"a", "b", "c"})
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	fileReader := bufio.NewReader(file)
	combos := make([]int, 0)
	for {
		line, e := fileReader.ReadBytes('\n')
		line = bytes.TrimRight(line, "\n")
		value, tf := strconv.Atoi(string(line))
		if tf == nil {
			combos = append(combos, value)
		} else {
			panic(tf)
		}
		if e == io.EOF {
			break
		}
	}
	// a := []string{"a", "b", "c"}
	comb := combinations.New(combos)
	comb.GenerateSet()
	out := comb.CheckSet()
	fmt.Println(out)

}
