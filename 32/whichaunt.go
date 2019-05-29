package main

import (
	"Advent2015/31/aunt"
	"bufio"
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
	fileReader := bufio.NewReader(file)
	aunts := make([]aunt.Aunt, 0)
	for {
		line, e := fileReader.ReadBytes('\n')
		removedNewline := strings.TrimRight(string(line), "\n")
		words := strings.Split(removedNewline, " ")
		id := string(words[1])
		id = strings.TrimRight(id, ":")
		intID, _ := strconv.Atoi(id)
		// fmt.Println(intID)
		thisAunt := aunt.New(intID)
		for i, v := range words {
			v = strings.Replace(v, ",", "", -1)
			v = strings.Replace(v, ":", "", -1)
			words[i] = v
		}
		i := 2
		for i < len(words) {
			key := words[i]
			stringValue := words[i+1]
			value, e := strconv.Atoi(stringValue)
			if e != nil {
				fmt.Println("String Value: ", stringValue)
				panic(e)
			}
			thisAunt.AddValue(key, value)
			i += 2
		}
		// fmt.Println(thisAunt)
		aunts = append(aunts, thisAunt)
		if e == io.EOF {
			break
		}
	}
	myAunt := aunt.New(501)
	myAunt.AddValue("children", 3)
	myAunt.AddValue("cats", 7)
	myAunt.AddValue("samoyeds", 2)
	myAunt.AddValue("pomeranians", 3)
	myAunt.AddValue("akitas", 0)
	myAunt.AddValue("vizslas", 0)
	myAunt.AddValue("goldfish", 5)
	myAunt.AddValue("trees", 3)
	myAunt.AddValue("cars", 2)
	myAunt.AddValue("perfumes", 1)

	// myMap := make(map[string]int, 0)
	// myMap["children"] = 3
	// myMap["cats"] = 7
	// myMap["samoyeds"] = 2
	// myMap["pomeranians"] = 3
	// myMap["akitas"] = 0
	// myMap["vizslas"] = 0
	// myMap["goldfish"] = 5
	// myMap["trees"] = 5
	// myMap["cars"] = 3
	// myMap["perfumes"] = 3

	for _, v := range aunts {
		out := v.Compare(myAunt)
		if out {
			fmt.Println(v.GetID())
		}
	}
}
