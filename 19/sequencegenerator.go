package main

import "fmt"

func main() {
	input := "3113322113"
	output := make([]byte, 0)
	for i, v := range input {
		letterCount := 1
		for _, w := range input[i:] {
			if w == v {
				letterCount++
			} else {
				break
			}
			fmt.Println("test")
		}
		byteCount := byte(letterCount)
		output = append(output, byteCount)
		fmt.Println(output, len(output), cap(output))
		// output = append(output, byte(v))
	}
	fmt.Println(string(output))
}
