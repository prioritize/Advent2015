package fileParser

import (
	"bufio"
	"bytes"
	"io"
	"os"
)

type fileReader struct {
	name         string
	replacements map[string]string
}

func New(fileName string) (map[string][]string, string) {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	var outString string
	fileReader := bufio.NewReader(file)
	replacements := make(map[string][]string, 0)
	for {
		line, e := fileReader.ReadBytes('\n')
		if len(line) == 1 {
			line, _ := fileReader.ReadBytes('\n')
			line = bytes.TrimRight(line, "\n")
			outString = string(line)
			break
		}
		line = bytes.TrimRight(line, "\n")
		words := bytes.Split(line, []byte(" "))

		elem := replacements[string(words[0])]
		if elem == nil {
			newElem := make([]string, 1)
			newElem[0] = string(words[2])
			replacements[string(words[0])] = newElem
		} else {
			elem = append(elem, string(words[2]))
			replacements[string(words[0])] = elem
		}
		if e == io.EOF {
			break
		}
	}
	return replacements, outString
}
