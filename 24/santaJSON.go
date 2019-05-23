package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

type Foo struct {
	X map[string]interface{}
}

// I took this from reddit.com/u/CremboC. First question I've had to turn to Reddit for assistance
func rec(f interface{}) (output float64) {
outer:
	switch fv := f.(type) {
	case []interface{}:
		for _, val := range fv {
			output += rec(val)
		}
	case float64:
		output += fv
	case map[string]interface{}:
		for _, val := range fv {
			if val == "red" {
				break outer
			}
		}
		for _, val := range fv {
			output += rec(val)
		}
	}

	return output
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	// b := []byte(`{"e":"green","c":141,"a":-11,"b":129,"d":"orange","f":"green"}`)
	fileReader := bufio.NewReader(file)
	data, _ := fileReader.ReadBytes('\n')
	// f := Foo{}
	var f interface{}
	e := json.Unmarshal(data, &f)
	var output float64
	output = rec(f)
	fmt.Println(output)
	if e != nil {
		panic(e)
	}
	// for k, v := range m {
	// 	switch vv := v.(type) {
	// 	case string:
	// 		fmt.Println(k, "is string", vv)
	// 	case float64:
	// 		fmt.Println(k, "is float64", vv)
	// 	case []interface{}:
	// 		fmt.Println(k, "is an array:")
	// 		for i, u := range vv {
	// 			fmt.Println(i, u)
	// 		}
	// 	default:
	// 		fmt.Println(k, "is of a type I don't know how to handle")
	// 	}
	// }
	// if err := json.Unmarshal(data, &f); err != nil {
	// 	panic(err)
	// }
	// if err := json.Unmarshal(data, &f.X); err != nil {
	// 	panic(err)
	// }
	// fmt.Println(f)
	// fmt.Println(len(f.X))
}
