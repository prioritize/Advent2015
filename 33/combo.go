package main

import (
	"Advent2015/33/combinations"
)

func main() {
	// comb := combinations.New([]string{"a", "b", "c"})
	a := []string{"a", "b", "c"}
	comb := combinations.New(a)
	comb.GenerateSet()

}
