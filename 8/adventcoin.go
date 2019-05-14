package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"strconv"
	"strings"
)

func main() {
	hash := md5.New()
	var prefix = "ckczppom"
	var suffix strings.Builder
	var hashOut []byte
	var index int
	for {
		// Reset the string builder to remove old values
		suffix.Reset()
		// Reset that hash function to remove old values
		hash.Reset()
		// Add the prefix provided to beginning of the suffix string
		suffix.WriteString(prefix)
		// Convert the current iteration index to a string
		stringNumber := strconv.Itoa(index)
		// Append the index value to the end of the suffix string
		suffix.WriteString(stringNumber)
		// Write the combined string into the md5hash
		io.WriteString(hash, suffix.String())
		// Calculate the md5Sum
		hashOut = hash.Sum(nil)
		// Write from the []byte value in hashOut to a []string
		stringHexValue := hex.EncodeToString(hashOut)
		// Set testValue equal to the first 5 values in stringHexValue
		testValue := stringHexValue[0:6]
		// Check if the first 5 values of hashOut are 0
		if testValue == "000000" {
			// Break out of the for loop when an appropriate value is found
			break
		}
		index++
	}
	// Print the appropriate value
	fmt.Println("The first value that results with 5 leading zeros is: ", index)
}
