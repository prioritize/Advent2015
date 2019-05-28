package main

import (
	"bufio"
	"io"
	"os"
	"strings"
)

type Person interface {
	AddValue(string, int)
	CheckValue(string) (int, bool)
	MakePerson(int)
}
type Sue struct {
	attributes map[string]int
	id         int
}

func (s Sue) MakePerson(id int) Sue {

	return s
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	fileReader := bufio.NewReader(file)
	line, e := fileReader.ReadBytes('\n')
	for {
		removedNewline := strings.TrimRight(string(line), "\n")
		id := removedNewline[1]

		if e == io.EOF {
			break
		}
	}
}
