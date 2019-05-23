package main

import "fmt"

// CheckDisallowed checks if any of the disallowed characters are in the input
func CheckDisallowed(input []byte) bool {
	// i, o and l are disallowed
	for _, v := range input {
		if v == 'i' || v == 'l' || v == 'o' {
			return true
		}
	}
	return false
}

// GenerateNext returns the next password
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

// ContainsStraight checks for 3 increasing characters in a row
func ContainsStraight(input []byte) bool {
	for i := 0; i < len(input)-2; i++ {

		if input[i] == (input[i+1]-1) && input[i] == (input[i+2]-2) {
			return true
		}
	}
	return false
}

// ContainsPairs checks for matching characters, returns when 2 are found
func ContainsPairs(input []byte) bool {
	pairCount := 0
	var firstPair byte
	for i := 0; i < len(input)-1; i++ {
		if input[i] == input[i+1] {
			if pairCount == 0 {
				firstPair = input[i]
				pairCount++
				i++
			}
			if pairCount == 1 {
				if firstPair == input[i] {
					continue
				} else {
					pairCount++
					i++
				}
			}
		}
	}
	if pairCount > 1 {
		return true
	}
	return false
}
func GenerateNextPassword(input []byte) []byte {
	for {
		GenerateNext(input)
		if !CheckDisallowed(input) {
			if ContainsPairs(input) {
				if ContainsStraight(input) {
					// fmt.Println(string(input))
					break
				}
			}
		}
	}
	return input
}
func main() {
	input := []byte("hepxcrrq")
	output := GenerateNextPassword(input)
	fmt.Println(string(output))
	output = GenerateNextPassword(output)
	fmt.Println(string(output))

}
