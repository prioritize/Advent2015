package main

import "fmt"

func CheckDisallowed(input []byte) bool {
	// i, o and l are disallowed
	for _, v := range input {
		if v == 'i' || v == 'l' || v == 'o' {
			return true
		}
	}
	return false
}
func GenerateNext(input []byte) []byte {
	// 8 runes, all lower case, rolling over from z->a
	rollover := true
	for i := len(input) - 1; i >= 0; i-- {
		if rollover {
			input[i]++
			rollover = false
		}
		if input[i] == 123 {
			input[i] = 97
			rollover = true
		} else {
			break
		}
	}
	return input
}
func ContainsStraight(input []byte) bool {
	return true
}
func ContainsPairs(input []byte) bool {
	return true
}
func HandleRollover(input []byte) {

}
func main() {
	input := []byte("hepxcrrq")
	for {

		GenerateNext(input)
		if !CheckDisallowed(input) {
			fmt.Println("valid")
		}
		fmt.Println("invalid")
	}

}
