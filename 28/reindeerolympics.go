package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Reindeer struct
type Reindeer struct {
	name     string
	speed    int
	duration int
	rest     int
	distance int
	period   int
	running  bool
	counter  int
	points   int
}

// Animal Interface implementation
type Animal interface {
	DistanceCovered()
	GetDistance() int
	AddPoint()
	GetPoints() int
	GetName() string
}

// CreateReindeer creates and returns a reindeer
func CreateReindeer(name string, speed, duration, rest int) Reindeer {
	var deer Reindeer
	deer.name = name
	deer.speed = speed
	deer.duration = duration
	deer.rest = rest
	deer.period = duration + rest
	deer.running = true
	return deer
}

// DistanceCovered is used on a per tick basis and calculates the distance covered
func (r *Reindeer) DistanceCovered() {
	r.counter++
	if r.running {
		r.distance += r.speed
		if r.counter >= r.duration {
			r.running = false
			r.counter = 0
		}
	} else {
		if r.counter >= r.rest {
			r.running = true
			r.counter = 0
		}
	}
}

// GetDistance returns distance of Reindeer
func (r Reindeer) GetDistance() int {
	return r.distance
}

// AddPoint adds a point to the reindeers total points
func (r *Reindeer) AddPoint() {
	r.points++
}

// GetPoints returns the point total for the reindeer
func (r *Reindeer) GetPoints() int {
	return r.points
}

// GetName gets the reindeers name
func (r *Reindeer) GetName() string {
	return r.name
}

// CalculateDistance calculates the distance a reindeer has traveled after a supplied distance
func CalculateDistance(deer Reindeer, duration int) int {
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
	team := make([]Animal, 0)
	for {
		newLine, e := fileReader.ReadBytes('\n')
		line := strings.Split(string(newLine), " ")
		name := line[0]
		speed, _ := strconv.Atoi(line[3])
		duration, _ := strconv.Atoi(line[6])
		rest, _ := strconv.Atoi(line[13])
		deer := CreateReindeer(name, speed, duration, rest)
		team = append(team, &deer)
		if e == io.EOF {
			break
		}
	}
	i := 0
	for i < 2503 {
		for _, v := range team {
			v.DistanceCovered()
		}
		max := 0
		for _, v := range team {
			if v.GetDistance() >= max {
				max = v.GetDistance()
			}
		}
		for _, v := range team {
			if v.GetDistance() == max {
				v.AddPoint()
			}
		}
		i++
	}
	for _, v := range team {
		fmt.Println(v.GetName(), v.GetPoints())
	}

}
