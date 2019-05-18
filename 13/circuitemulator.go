package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Operation is an interface for operations that can be fed in through the input text file
type Operation interface {
	SetInputs(map[string]string)
	CheckInputs() bool
	RShift()
	Or()
	LShift()
	And()
	Not()
	Assign()
}

// Gate is struct that implements a node that takes two inputs and an operation and provides an output
type Gate struct {
	inputs     [2]uint16
	mapInputs  [2]string
	valuesSet  [2]bool
	output     [1]uint16
	mapOutput  [1]string
	outputBool [1]bool
	op         [1]string
}

//Assign  is a struct that implements an node that takes one input and an operation and provides an output
type Assign struct {
	inputs     [1]uint16
	mapInputs  [1]string
	valuesSet  [1]bool
	output     [1]uint16
	mapOutput  [1]string
	outputBool [1]bool
	op         [1]string
}

//RShift needs two inputs, a value and a offset and provides an output
func (g Gate) RShift() {
	if g.outputBool[0] == false {
		g.output[0] = g.inputs[0] >> g.inputs[1]
		g.outputBool[0] = true
	}
}

//LShift needs two inputs, a value and a offset and provides an output
func (g Gate) LShift() {
	if g.outputBool[0] == false {
		g.output[0] = g.inputs[0] << g.inputs[1]
		g.outputBool[0] = true
	}
}

// And needs two inputs, two values and provides an output
func (g Gate) And() {
	if g.outputBool[0] == false {
		g.output[0] = g.inputs[0] & g.inputs[1]
		g.outputBool[0] = true
	}
}

// Or needs two inputs, two values and provides an output
func (g Gate) Or() {
	if g.outputBool[0] == false {
		g.output[0] = g.inputs[0] | g.inputs[1]
		g.outputBool[0] = true
	}
}

// Not is not implemented as it is not required
func (g Gate) Not() {
	// Not Implemented. Not is an Assign function
}

// Assign is not implemented as it is not required
func (g Gate) Assign() {
	// Not implemented. Assign is an Assign function
}

// CheckInputs loops through the valuesSet slice and checks for any falses. If false, return false, else return true
func (g Gate) CheckInputs() bool {
	// TODO: Loop through the valuesSet bool, if all all are true, return true
	// TODO: Can perform some error checking here and check the values in the map to ensure no errors  were made
	for _, v := range g.valuesSet {
		if v == false {
			return false
		}
		return true
	}
	return true
}

// RShift is not implemented as it is not required
func (a Assign) RShift() {
	//Not implemented. RShift is a Gate function
}

// LShift is not implemented as it is not required
func (a Assign) LShift() {
	// Not implemented. LShift is a Gate function
}

// And is not implemented as it is not required
func (a Assign) And() {
	// Not implemented. LShift is a Gate function
}

// Or is not implemented as it is not required
func (a Assign) Or() {
	// Not implemented. LShift is a Gate function
}

// Not requires one input. It provides one output
func (a Assign) Not() {
	if a.outputBool[0] == false {
		a.output[0] = ^a.inputs[0]
		a.outputBool[0] = true
	}
}

// Assign requires one input. It provides one output
func (a Assign) Assign() {
	if a.outputBool[0] == false {
		a.output[0] = a.inputs[0]
		a.outputBool[0] = true
	}
}

func (a Assign) CheckInputs() bool {
	// TODO: Loop through the valuesSet bool, if all all are true, return true
	// TODO: Can perform some error checking here and check the values in the map to ensure no errors  were made
	return true
}

// SetInputs checks the values in the map and if they are integers places them into the input values in the RShift object
func (g Gate) SetInputs(m map[string]string) {
	// Get the values from the map to check if they are integers
	for i, v := range g.mapInputs {
		elem := m[v]
		mapValue, e := strconv.Atoi(elem)
		if (e) == nil {
			g.inputs[i] = uint16(mapValue)
		}
	}
}

// SetInputs checks the values in the map and if they are integers places them into the input values in the RShift object
func (a Assign) SetInputs(m map[string]string) {
	// Get the values from the map to check if they are integers
	for i, v := range a.mapInputs {
		elem := m[v]
		mapValue, e := strconv.Atoi(elem)
		if e == nil {
			a.inputs[i] = uint16(mapValue)
		}
	}
}

// MakeGate processes a line and returns a Gate
func MakeGate(line []string) Gate {
	var g Gate
	g.mapInputs[0] = line[0]
	g.mapInputs[1] = line[2]
	g.op[0] = line[1]
	g.mapOutput[0] = line[4]
	return g
}

// MakeAssign processes a line and returns a Gate
func MakeAssign(line []string) Assign {
	// TODO: finish this implementation
	var a Assign
	return a
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	// objectCommand := make(map[string]Operation)
	nodeSlice := make([]Operation, 0)
	commands := make(map[string]Operation)
	nodes := make(map[string]int)
	bufferedFile := bufio.NewReader(file)
	index := 0
	for {
		line, e := bufferedFile.ReadBytes('\n')
		if e == io.EOF {
			break
		}
		trimmedLine := strings.TrimRight(string(line), "\n")
		splitLine := strings.Split(trimmedLine, " ")

		switch {
		case splitLine[1] == "OR" || splitLine[1] == "AND" || splitLine[1] == "RSHIFT" || splitLine[1] == "LSHIFT":
			nodeSlice = append(nodeSlice, MakeGate(splitLine))
		case splitLine[0] == "NOT" || splitLine[1] == "->":
			nodeSlice = append(nodeSlice, MakeAssign(splitLine))
		}
		// TODO: Start looping through nodeSlice to build the objects

		elem := commands[strconv.Itoa(index)]
		fmt.Printf("Element %d: %s\n", index, elem)
		index++
	}
	fmt.Println(nodeSlice)
	elem := commands["0"]
	fmt.Println(nodes["a"])
	fmt.Println(elem)
}
