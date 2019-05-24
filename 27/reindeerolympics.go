package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Reindeer struct {
	name     string
	speed    int
	duration int
	rest     int
	distance int
	period   int
}

func CreateReindeer(name string, speed, duration, rest int) Reindeer {
	var deer Reindeer
	deer.name = name
	deer.speed = speed
	deer.duration = duration
	deer.rest = rest
	deer.period = duration + rest
	return deer
}
func CalculcateDistance(deer Reindeer, duration int) int {
	var cycles int
	var distance int
	cycles = duration / deer.period
	distance = cycles * deer.duration * deer.speed
	secondsLeft := duration - deer.period*cycles
	if secondsLeft >= deer.duration {
		distance += deer.speed * deer.duration
	} else {
		distance += deer.speed * secondsLeft
	}
	return distance
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	fileReader := bufio.NewReader(file)
	team := make([]Reindeer, 0)
	for {
		newLine, e := fileReader.ReadBytes('\n')
		line := strings.Split(string(newLine), " ")
		name := line[0]
		speed, _ := strconv.Atoi(line[3])
		duration, _ := strconv.Atoi(line[6])
		rest, _ := strconv.Atoi(line[13])
		team = append(team, CreateReindeer(name, speed, duration, rest))
		if e == io.EOF {
			break
		}
	}
	for _, v := range team {
		fmt.Println(CalculcateDistance(v, 2503))
	}

}
