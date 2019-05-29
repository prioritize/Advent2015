package combinations

import (
	"fmt"
	"math"
)

type Combination struct {
	set     [][]int
	initial []int
}

func New(initial []int) Combination {
	var c Combination
	c.initial = initial
	c.set = make([][]int, 0)
	for i := range c.set {
		c.set[i] = make([]int, 0)
	}
	return c
}
func (c *Combination) GenerateSet() {
	combos := uint(math.Pow(2, float64(len(c.initial))))
	var i uint
	for i < combos {
		currentCombo := make([]int, 0)
		for j, x := range c.initial {
			if i&(1<<uint(j)) != 0 {
				currentCombo = append(currentCombo, x)
			}
		}
		c.set = append(c.set, currentCombo)
		i++
	}
}
func (c Combination) CheckSet() int {
	matches := 0
	minCount := 0
	for _, v := range c.set {
		total := 0
		for _, j := range v {
			total += j
		}
		if total == 150 {
			matches++
			if len(v) == 4 {
				minCount++
			}
		}
	}
	fmt.Println(minCount)
	return matches
}
