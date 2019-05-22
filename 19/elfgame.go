package main

import (
	"fmt"
	"strconv"
	"strings"
)

// GetConsecutives iterates through a slice checking for continuous values
func GetConsecutives(input string, current byte) int {
	output := 0
	for _, v := range input {
		if v == rune(current) {
			output++
		} else {
			return output
		}
	}
	return output
}
func main() {
	var b strings.Builder
	// input := "121211"
	// input := "1"
	input := "3113322113"
	offset := 0
	// Run through the string until one less than the length, this enables checking of the last value
	loops := 0
	for {
		for {
			offset++
			holder := input[offset:]
			consecutives := GetConsecutives(holder, input[offset-1])
			if consecutives == 0 {
				b.WriteByte([]byte(strconv.Itoa(1))[0])
				b.WriteByte(input[offset-1])
			} else {
				b.WriteByte([]byte(strconv.Itoa(consecutives + 1))[0])
				b.WriteByte(input[offset-1])
			}
			offset += consecutives
			if offset >= len(input) {
				break
			}
		}
		loops++
		fmt.Println(b.String())
		offset = 0
		input = b.String()
		if loops == 40 {
			break
		}
		b.Reset()
	}
	fmt.Println(len(b.String()))

}
