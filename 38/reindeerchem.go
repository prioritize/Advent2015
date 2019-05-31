package main

import (
	fileParser "Advent2015/37/reader"
	"fmt"
	"os"
	"strings"
)

func main() {
	args := os.Args
	fmt.Println(args)
	var indices []int
	outMap, outString := fileParser.New(args[1])
	fmt.Println(outMap, outString)
	// testString := "HOH"
	// fmt.Println(strings.Count(outString, "Al"))
	output := make(map[string]string, 0)
	for k, v := range outMap {
		fmt.Println(k, strings.Count(outString, k))
		indices = GetIndices(outString, k)
		for _, x := range v {
			for _, w := range indices {
				output[ReplaceSubString(outString, k, x, w)] = ""
			}
		}
	}
	fmt.Println(len(output))
}
func GetIndices(s, sub string) []int {
	indices := make([]int, 0)
	for i, v := range s {
		if v == rune(sub[0]) {
			//First character matches
			if sub == s[i:i+len(sub)] {
				indices = append(indices, i)
			}
		}
	}

	lenInt := strings.Count(s, sub)
	if lenInt != len(indices) {
		panic("This wrong")
	}
	return indices
}
func ReplaceSubString(s, sub, new string, index int) string {
	subLen := len(sub)
	var pre []byte

	// copy(pre, []byte(s[0:index+1]))
	for _, v := range []byte(s[0:index]) {
		pre = append(pre, v)
	}
	for _, v := range new {
		pre = append(pre, byte(v))
	}
	for _, v := range []byte(s[(index + subLen):]) {
		pre = append(pre, v)
	}
	return string(pre)
}
