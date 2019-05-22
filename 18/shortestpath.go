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

// Location defines functions to be used in City
type Location interface {
	GetDistance() (int, bool)
	GetName() string
	Map()
}

// City holds the values for a City
type City struct {
	name           string
	destinationMap map[string]int
}

// GetName returns the name of the city
func (c City) GetName() string {
	return c.name
}

// GetDistance returns the distance between the city object and a city name
func (c City) GetDistance(city string) (int, bool) {
	elem, e := c.destinationMap[city]
	if e == false {
		fmt.Println("This shouldn't be happening", c.name, city)
	}
	return elem, e

}

//LocationExists determines if a location exists in a slice of cities
func LocationExists(locations []City, specific string) bool {
	for _, v := range locations {
		if v.GetName() == specific {
			// TODO: Consider setting the keys here, this would mean destination and distance
			return true
		}
	}
	return false
}

// DestinationExists determines if a city is listed in a cities dictionary
func DestinationExists(city City, destination string) bool {
	for k := range city.destinationMap {
		if k == destination {
			return true
		}
	}
	return false
}

// GetCity returns a city from a slice of cities
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

// CreateCity creates a new City object and instantiates the default values
func CreateCity(name string) City {
	var newCity City
	newCity.name = name
	newCity.destinationMap = make(map[string]int, 0)
	// newCity.routes = make(map[string]string, 0)

	return newCity
}

// AddDestination addes a destination and distance to a city
func AddDestination(thisCity City, name string, distance string) City {
	value, err := strconv.Atoi(distance)
	if err != nil {
		panic(err)
	}
	thisCity.destinationMap[name] = value
	return thisCity
}

// Permute generates all the possible permutations give a supplied string. Return will be !len([]string)
func Permute(in []string) [][]string {
	c := make([]int, len(in))
	out := make([][]string, 0)
	iteration := make([]string, len(in))
	copy(iteration, in)

	for i := range out {
		out[i] = make([]string, 0)
	}
	out = append(out, iteration)
	i := 0
	for i < len(iteration) {
		test := make([]string, len(in))
		if c[i] < i {
			if i%2 == 0 {
				iteration[0], iteration[i] = iteration[i], iteration[0]
			} else {
				iteration[c[i]], iteration[i] = iteration[i], iteration[c[i]]
			}
			// out = append(out, iteration)
			c[i]++
			i = 0
			copy(test, iteration)
			out = append(out, test)
		} else {
			c[i] = 0
			i++
		}
	}
	return out
}

// CheckMultiples checks if a city exists in a slice
func CheckMultiples(cities []string, city string) bool {
	for _, v := range cities {
		if v == city {
			return true
		}
	}
	return false
}
func main() {
	file, _ := os.Open("input.txt")
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
	runningTotal := 0
	min := 1000000
	max := 0
	for _, v := range out {
		for j := 0; j < len(v)-1; j++ {
			city1, _, _ := GetCity(cities, v[j])
			interim, _ := city1.GetDistance(v[j+1])
			runningTotal += interim
		}
		if runningTotal < min {
			min = runningTotal
		}
		if runningTotal > max {
			max = runningTotal
		}
		// fmt.Println(runningTotal)
		// fmt.Println(min)
		// fmt.Println(max)
		runningTotal = 0
	}
	fmt.Println(min)
	fmt.Println(max)

}
