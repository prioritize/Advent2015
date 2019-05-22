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

type Location interface {
	GetDistance() (int, bool)
	GetName() string
	Map()
}
type City struct {
	name           string
	destinationMap map[string]int
}

func (c City) GetName() string {
	return c.name
}
func (c City) GetDistance(city string) (int, bool) {
	elem, e := c.destinationMap[city]
	if e == false {
		fmt.Println("This shouldn't be happening", c.name, city)
	}
	return elem, e

}
func LocationExists(locations []City, specific string) bool {
	for _, v := range locations {
		if v.GetName() == specific {
			// TODO: Consider setting the keys here, this would mean destination and distance
			return true
		}
	}
	return false
}
func DestinationExists(city City, destination string) bool {
	for k, _ := range city.destinationMap {
		if k == destination {
			return true
		}
	}
	return false
}
func GetCity(cities []City, name string) (City, bool, int) {
	var outCity City
	var index int
	for i, v := range cities {
		if v.GetName() == name {
			outCity = v
			return outCity, true, i
		}
	}
	return outCity, false, index
}
func CreateCity(name string) City {
	var newCity City
	newCity.name = name
	newCity.destinationMap = make(map[string]int, 0)
	// newCity.routes = make(map[string]string, 0)

	return newCity
}
func AddDestination(thisCity City, name string, distance string) City {
	value, err := strconv.Atoi(distance)
	if err != nil {
		panic(err)
	}
	thisCity.destinationMap[name] = value
	return thisCity
}
func Permute(in []string) [][]string {
	c := make([]int, len(in))
	out := make([][]string, 0)
	for i := range out {
		out[i] = make([]string, 0)
	}
	fmt.Println(in)
	out = append(out, in)
	i := 0
	for i < len(in) {
		if c[i] < i {
			if i%2 == 0 {
				in[0], in[i] = in[i], in[0]
			} else {
				in[c[i]], in[i] = in[i], in[c[i]]
			}
			fmt.Println(in)
			out = append(out, in)
			c[i]++
			i = 0
		} else {
			c[i] = 0
			i++
		}
	}
	return out
}
func CheckMultiples(cities []string, city string) bool {
	for _, v := range cities {
		if v == city {
			return true
		}
	}
	return false
}
func main() {
	file, _ := os.Open("test_input.txt")
	fileReader := bufio.NewReader(file)
	maps := make([][]byte, 0)
	cityNames := make([]string, 0)
	for i := range maps {
		maps[i] = make([]byte, 0)
	}
	for {
		line, e := fileReader.ReadBytes('\n')
		line = bytes.Replace(line, []byte("\n"), []byte(""), -1)
		if e == io.EOF {
			break
		}
		maps = append(maps, line)
	}
	cities := make([]City, 0)
	end := 2
	d := 4
	start := 0
	// var originCity, destinationCity City
	for _, v := range maps {
		input := strings.Split(string(v), " ")
		origin := input[start]
		fmt.Println(input)
		destination := input[end]
		distance := input[d]
		locations := make([]string, 0)
		locations = append(locations, origin)
		locations = append(locations, destination)

		// distance := input[d]
		for _, v := range locations {
			if !LocationExists(cities, v) {
				originCity := CreateCity(v)
				cities = append(cities, originCity)
			}
		}
		city, _, i := GetCity(cities, origin)
		city = AddDestination(city, destination, distance)
		cities[i] = city
		city, _, i = GetCity(cities, destination)
		city = AddDestination(city, origin, distance)
		cities[i] = city

	}
	for _, v := range cities {
		cityNames = append(cityNames, v.name)
	}
	out := Permute(cityNames)
	for _, v := range out {
		fmt.Println(v)
	}
	runningTotal := 0
	var path []string
	min := 1000000
	for _, v := range out {
		for j := 0; j < len(v)-1; j++ {
			city1, _, _ := GetCity(cities, v[j])
			interim, _ := city1.GetDistance(v[j+1])
			// fmt.Println(city1.name, v[j+1], "Interim: ", interim)
			runningTotal += interim
		}
		if runningTotal < min {
			min = runningTotal
			path = v
		}
		// fmt.Println(runningTotal)
		runningTotal = 0
	}

	fmt.Println(path)
	// fmt.Println(min)
	for _, v := range cities {
		fmt.Println(v)
	}
	str1 := "London"
	str2 := "Dublin"
	str3 := "Belfast"

	var test1 []string
	var test2 []string

	test1 = append(test1, str3)
	test1 = append(test1, str2)
	test1 = append(test1, str1)

	test2 = append(test2, str1)
	test2 = append(test2, str2)
	test2 = append(test2, str3)

	fmt.Println(test1)
	fmt.Println(test2)
	var test3 [][]string
	test3 = append(test3, test1)
	test3 = append(test3, test2)
	fmt.Println(test3)

}
