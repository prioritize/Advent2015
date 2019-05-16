package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Operation interface {
	BitOperation() uint16
}
type RShift struct {
	input, offset, output uint16
}
type LShift struct {
	input, offset, output uint16
}
type Not struct {
	input, output uint16
}
type Assign struct {
	input, output uint16
}
type And struct {
	input1, input2, output uint16
}
type Or struct {
	input1, input2, output uint16
}

func (r RShift) BitOperation() uint16 {
	r.output = r.input >> r.offset
	return r.output
}
func (o Or) BitOperation() uint16 {
	o.output = o.input1 | o.input2
	return o.output
}
func (l LShift) BitOperation() uint16 {
	l.output = l.input << l.offset
	return l.output
}
func (a And) BitOperation() uint16 {
	a.output = a.input1 | a.input2
	return a.output
}
func (n Not) BitOperation() uint16 {
	n.output = ^n.input
	return n.output
}

func FileToMap() map[string]Operation {
	newMap := make(map[string]Operation)

	return newMap
}

// func DetermineOperation(line []string, op Operation) Operation {
// 	operations := [5]string{"NOT", "RSHIFT", "OR", "AND", "LSHIFT"}
// 	for _, v := range line {
// 		for _, w := range operations {
// 			if v == w {
// 				op.op = v
// 				// fmt.Println(op.op)
// 				break
// 			}
// 		}
// 	}
// 	if len(line) == 3 {
// 		op.op = "ASSIGN"
// 	}
// 	return op
// }
// func ParseLine(line string) Operation {
// 	op := Operation{}
// 	splitLine := strings.Split(line, " ")
// 	op = DetermineOperation(splitLine, op)
// 	switch {
// 	case op.op == "NOT":
// 		op.inOne = splitLine[1]
// 		op.out = splitLine[3]
// 	case op.op == "AND" || op.op == "OR":
// 		op.inOne = splitLine[0]
// 		op.inTwo = splitLine[2]
// 		op.out = splitLine[4]
// 	case op.op == "RSHIFT" || op.op == "LSHIFT":
// 		op.inOne = splitLine[0]
// 		op.inTwo = splitLine[2]
// 		op.out = splitLine[4]
// 	case op.op == "ASSIGN":
// 		op.inOne = splitLine[0]
// 		op.out = splitLine[2]
// 	}
// 	return op
// }

// func GenerateNodes(commands map[string]Operation) map[string]int {
// 	nodes := make(map[string]int)
// 	for _, v := range commands {
// 		_, e0 := strconv.Atoi(v.inOne)
// 		if e0 != nil {
// 			nodes[v.inOne] = -1
// 		}
// 		_, e1 := strconv.Atoi(v.inTwo)
// 		if e1 != nil {
// 			nodes[v.inTwo] = -1
// 		}
// 		_, e2 := strconv.Atoi(v.out)
// 		if e2 != nil {
// 			nodes[v.out] = -1
// 		}
// 	}
// 	return nodes
// }
// func EvaluateNodes(commands map[string]Operation, nodes map[string]int) (map[string]Operation, map[string]int) {
// 	for _, v := range commands {
// 		switch {
// 		case v.op == "RSHIFT":
// 			value, err := strconv.Atoi(v.inOne)
// 			if err == nil {

// 			}
// 		case v.op == "LSHIFT":
// 			value, err := strconv.Atoi(v.inOne)
// 			if err == nil {

// 			}
// 		case v.op == "NOT":
// 			value, err := strconv.Atoi(v.inOne)
// 			if err == nil {

// 			}
// 		case v.op == "ASSIGN":
// 			value, err := strconv.Atoi(v.inOne)
// 			if err == nil {

// 			}
// 		}

// 		if err == nil {
// 			// The values are available
// 		}
// 	}

// 	return commands, nodes
// }

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	objectCommand := make(map[string]Operation)
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
		commands[strconv.Itoa(index)] = ParseLine(trimmedLine)
		elem := commands[strconv.Itoa(index)]
		fmt.Printf("Element %d: %s\n", index, elem)
		index++
	}
	// nodes = GenerateNodes(commands)
	fmt.Println(nodes)
	elem := commands["0"]
	fmt.Println(nodes["a"])
	fmt.Println(elem)
}
