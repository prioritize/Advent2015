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

type Person struct {
	name      string
	happiness map[string]int
}

func CreatePerson(name string) Person {
	var p Person
	p.happiness = make(map[string]int, 0)
	p.name = name
	return p
}
func AddNeighbor(p Person, neighbor string, value int) {
	p.happiness[neighbor] = value
}
func FindPerson(p []Person, name string) (int, bool) {
	for i, v := range p {
		if v.name == name {
			return i, true
		}
	}
	return len(p) + 1, false
}
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

func CalculateHappiness(names []string, people []Person) int {
	total := 0
	var neighbor1 string
	var neighbor2 string
	for i := 0; i < len(names); i++ {
		switch {
		case i == 0:
			neighbor1 = names[len(names)-1]
			neighbor2 = names[i+1]
		case i == len(names)-1:
			neighbor1 = names[0]
			neighbor2 = names[i-1]
		default:
			neighbor1 = names[i-1]
			neighbor2 = names[i+1]
		}
		index, _ := FindPerson(people, names[i])
		total += people[index].happiness[neighbor1]
		total += people[index].happiness[neighbor2]
	}
	return total
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	people := make([]Person, 0)
	fileReader := bufio.NewReader(file)
	for {
		line, e := fileReader.ReadBytes('\n')
		line = bytes.TrimRight(line, "\n")
		line = bytes.TrimRight(line, ".")
		split := strings.Split(string(line), " ")
		person := split[0]
		sign := split[2]
		value := split[3]
		neighbor := split[10]
		index, exists := FindPerson(people, string(person))
		var thisPerson Person
		if !exists {
			thisPerson = CreatePerson(string(person))
			people = append(people, thisPerson)
		} else {
			thisPerson = people[index]
		}
		dislike, _ := strconv.Atoi(value)
		if sign == string("lose") {
			dislike = -1 * dislike
		}
		AddNeighbor(thisPerson, string(neighbor), dislike)
		people = append(people)

		if e == io.EOF {
			break
		}
	}
	var names []string
	for _, v := range people {
		names = append(names, v.name)
	}
	permutations := Permute(names)
	min := 1000000
	max := -1000000
	for _, v := range permutations {
		current := CalculateHappiness(v, people)
		if current > max {
			max = current
		}
		if current < min {
			min = current
		}
	}
	fmt.Println("Min: ", min)
	fmt.Println("Max: ", max)

}
