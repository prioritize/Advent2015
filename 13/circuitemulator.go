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
	SetInputs(map[string]string) Operation
	CheckInputs(m map[string]string) (Operation, bool)
	RShift(m map[string]string) Operation
	Or(m map[string]string) Operation
	LShift(m map[string]string) Operation
	And(m map[string]string) Operation
	Not(m map[string]string) Operation
	Assign(m map[string]string) Operation
	GetOp() string
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
func (g Gate) RShift(m map[string]string) Operation {
	if g.outputBool[0] == false {
		g.output[0] = g.inputs[0] >> g.inputs[1]
		g.outputBool[0] = true
	}
	m[g.mapOutput[0]] = strconv.Itoa(int(g.output[0]))
	return g
}

//LShift needs two inputs, a value and a offset and provides an output
func (g Gate) LShift(m map[string]string) Operation {
	if g.outputBool[0] == false {
		g.output[0] = g.inputs[0] << g.inputs[1]
		g.outputBool[0] = true
	}
	m[g.mapOutput[0]] = strconv.Itoa(int(g.output[0]))
	return g
}

// And needs two inputs, two values and provides an output
func (g Gate) And(m map[string]string) Operation {
	if g.outputBool[0] == false {
		g.output[0] = g.inputs[0] & g.inputs[1]
		g.outputBool[0] = true
	}
	m[g.mapOutput[0]] = strconv.Itoa(int(g.output[0]))
	return g
}

// Or needs two inputs, two values and provides an output
func (g Gate) Or(m map[string]string) Operation {
	if g.outputBool[0] == false {
		g.output[0] = g.inputs[0] | g.inputs[1]
		g.outputBool[0] = true
	}
	m[g.mapOutput[0]] = strconv.Itoa(int(g.output[0]))
	return g
}

// Not is not implemented as it is not required
func (g Gate) Not(m map[string]string) Operation {
	// Not Implemented. Not is an Assign function
	return g
}

// Assign is not implemented as it is not required
func (g Gate) Assign(m map[string]string) Operation {
	// Not implemented. Assign is an Assign function
	return g
}

// CheckInputs loops through the valuesSet slice and checks for any falses. If false, return false, else return true
func (g Gate) CheckInputs(m map[string]string) (Operation, bool) {
	// TODO: Loop through the valuesSet bool, if all all are true, return true
	// TODO: Can perform some error checking here and check the values in the map to ensure no errors  were made
	var b Operation
	for _, v := range g.valuesSet {
		if v == false {
			return g, false
		}
	}
	switch {
	case g.op[0] == "RSHIFT":
		b = g.RShift(m)
	case g.op[0] == "LSHIFT":
		b = g.LShift(m)
	case g.op[0] == "AND":
		b = g.And(m)
	case g.op[0] == "OR":
		b = g.Or(m)
	}
	return b, true
}

// RShift is not implemented as it is not required
func (a Assign) RShift(m map[string]string) Operation {
	//Not implemented. RShift is a Gate function
	return a
}

// LShift is not implemented as it is not required
func (a Assign) LShift(m map[string]string) Operation {
	// Not implemented. LShift is a Gate function
	return a
}

// And is not implemented as it is not required
func (a Assign) And(m map[string]string) Operation {
	// Not implemented. LShift is a Gate function
	return a
}

// Or is not implemented as it is not required
func (a Assign) Or(m map[string]string) Operation {
	// Not implemented. LShift is a Gate function
	return a
}

// Not requires one input. It provides one output
func (a Assign) Not(m map[string]string) Operation {
	if a.outputBool[0] == false {
		a.output[0] = ^a.inputs[0]
		a.outputBool[0] = true
	}
	m[a.mapOutput[0]] = strconv.Itoa(int(a.output[0]))
	return a
}

// Assign requires one input. It provides one output
func (a Assign) Assign(m map[string]string) Operation {
	if a.outputBool[0] == false {
		a.output[0] = a.inputs[0]
		a.outputBool[0] = true
	}
	m[a.mapOutput[0]] = strconv.Itoa(int(a.output[0]))
	return a
}

// GetOp returns g.op[0]
func (g Gate) GetOp() string {
	return g.op[0]
}

// GetOp returns g.op[0]
func (a Assign) GetOp() string {
	return a.op[0]
}

// CheckInputs loops through all inputBool values (valuesSet) and if all return true returns true
func (a Assign) CheckInputs(m map[string]string) (Operation, bool) {
	// TODO: Loop through the valuesSet bool, if all all are true, return true
	// TODO: Can perform some error checking here and check the values in the map to ensure no errors  were made
	var b Operation
	for _, v := range a.valuesSet {
		if v == false {
			return a, false
		}
	}
	// TODO: Consider executing the operation here and copying the calculated value into it's location in the dictionary
	switch {
	case a.op[0] == "ASSIGN":
		b = a.Assign(m)
	case a.op[0] == "NOT":
		b = a.Not(m)
	}
	return b, true
}

// SetInputs checks the values in the map and if they are integers places them into the input values in the RShift object
func (g Gate) SetInputs(m map[string]string) Operation {
	// Get the values from the map to check if they are integers
	for i, v := range g.mapInputs {
		elem := m[v]
		mapValue, e := strconv.Atoi(elem)
		if (e) == nil {
			g.inputs[i] = uint16(mapValue)
			g.valuesSet[i] = true
		}
	}
	return g
}

// SetInputs checks the values in the map and if they are integers places them into the input values in the RShift object
func (a Assign) SetInputs(m map[string]string) Operation {
	// Get the values from the map to check if they are integers
	for i, v := range a.mapInputs {
		elem := m[v]
		mapValue, e := strconv.Atoi(elem)
		if e == nil {
			a.inputs[i] = uint16(mapValue)
			a.valuesSet[i] = true
		}
	}
	return a
}

// MakeGate processes a line and returns a Gate
func MakeGate(line []string) Gate {
	var g Gate
	g.mapInputs[0] = line[0]
	g.mapInputs[1] = line[2]
	// TODO: Possibly call SetInputs from this location
	// TODO: Would require changing SetInputs to return a Gate instead of a map
	// TODO: May make a large amount of sense to not remove large value passes
	for i, v := range g.mapInputs {
		w, err := strconv.Atoi(v)
		if err == nil {
			g.inputs[i] = uint16(w)
			g.valuesSet[i] = true
		}
	}
	g.op[0] = line[1]
	g.mapOutput[0] = line[4]
	return g
}

// MakeAssign processes a line and returns a Gate
func MakeAssign(line []string) Assign {
	// TODO: Possibly call SetInputs from this location
	// TODO: Would require changing SetInputs to return a Gate instead of a map
	// TODO: May make a large amount of sense to not remove large value passes
	var a Assign
	if line[0] == "NOT" {
		a.mapInputs[0] = line[1]
		a.mapOutput[0] = line[3]
		a.op[0] = line[0]
	} else {
		a.mapInputs[0] = line[0]
		a.mapOutput[0] = line[2]
		a.op[0] = "ASSIGN"
	}
	for i, v := range a.mapInputs {
		w, err := strconv.Atoi(v)
		if err == nil {
			a.inputs[i] = uint16(w)
			a.valuesSet[i] = true
		}
	}
	return a
}

// BuildMap builds a map from a provided []string and returns a map[string]string. The file needs to be in the format provided by Advent Of Code
func BuildMap(line []string, m map[string]string) map[string]string {
	for _, v := range line {
		if v == "RSHIFT" || v == "LSHIFT" || v == "OR" || v == "AND" || v == "NOT" || v == "->" {
			continue
		} else {
			_, err := strconv.Atoi(v)
			if err != nil {
				m[v] = ""
			}
		}
	}
	return m
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	// objectCommand := make(map[string]Operation)
	nodeSlice := make([]Operation, 0)
	nodeValues := make(map[string]string, 0)
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
		nodeValues = BuildMap(splitLine, nodeValues)
		index++
	}

	// TODO: Return nodes from all the Gate and Assign functions to enable storing of the values in the object slice
	// TODO: Add function that checks for truth in the outputBool, if true, set the output value into the key value in the map

	for x := 0; x < 1000; x++ {
		for i, v := range nodeSlice {
			v = v.SetInputs(nodeValues)
			node, check := v.CheckInputs(nodeValues)
			if check == true {
				nodeSlice[i] = node
				// TODO: Write a function get to the output node OR handle the setting of the key value in a separate function which will require
				// TODO: passing of the nodeValues map
			}
		}
	}
	// TODO: Build the dictionary of objects
	// TODO: Parse through them and start placing values
	// fmt.Println(nodeSlice)
	fmt.Println(nodeValues)
	fmt.Println(len(nodeValues))
}
