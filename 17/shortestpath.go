package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
)

type Location interface {
	CalculateDistance()
	GetName() string
}
type City struct {
	name         string
	distanceList map[string]int
}

func (c City) GetName() string {
	return c.name
}
func (c City) CalculateDistance() {

}
func LocationExists(locations []City, specific string) bool {
	for _, v := range locations {
		if v.GetName() == specific {
			// TODO: Consider setting the keys here, this would mean destination and distance
		}
	}
	return false
}
func CreateCity(name string) City {
	var newCity City
	newCity.name = name
	return newCity

}

func main() {
	file, _ := os.Open("input.txt")
	fileReader := bufio.NewReader(file)
	maps := make([][]byte, 0)
	for i := range maps {
		maps[i] = make([]byte, 0)
	}
	for {
		line, e := fileReader.ReadBytes('\n')
		if e == io.EOF {
			break
		}
		line = bytes.Replace(line, []byte("\n"), []byte(""), -1)
		maps = append(maps, line)
	}
	locationSlice := make([]City, 0)
	for _, v := range maps {
		stringInput := strings.Split(string(v), " ")
		if !LocationExists(locationSlice, stringInput[0]) {
			newCity := CreateCity(string(stringInput[0]))
			locationSlice = append(locationSlice, newCity)
		}
	}
	fmt.Println(locationSlice)
}
