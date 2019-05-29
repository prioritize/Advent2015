package combinations

import (
	"fmt"
	"math"
)

type Combination struct {
	set     [][]string
	initial []string
}

func New(initial []string) Combination {
	var c Combination
	c.initial = initial
	c.set = make([][]string, 0)
	for i := range c.set {
		c.set[i] = make([]string, 0)
	}
	return c
}
func (c *Combination) GenerateSet() {
	combos := uint(math.Pow(2, float64(len(c.initial))))
	fmt.Println(combos)
	var i uint
	for i < combos {
		currentCombo := make([]string, 0)
		for j, x := range c.initial {
			if i&(1<<uint(j)) != 0 {
				currentCombo = append(currentCombo, x)
			}
		}
		c.set = append(c.set, currentCombo)
		i++
	}
	fmt.Println(c.set)
}
