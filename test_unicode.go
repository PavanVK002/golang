package main

import (
	"fmt"
	"strconv"
	"unicode"
)

func main() {
	// Given string
	str := "What is 5678 plus 13?"

	// Variables to store extracted numbers
	var numbers []int

	// Variables to accumulate digits forming a number
	var numStr string

	// Iterate over the characters of the string
	for _, char := range str {
		// Check if the character is a digit
		// fmt.Printf("%s\n",char)
		if unicode.IsDigit(char) {
			// If it's a digit, add it to the current number string
			numStr += string(char)
			// fmt.Println(numStr)

		} else {
			// If it's not a digit, check if we have a number string to convert
			if numStr != "" {
				// Convert the number string to an integer and append it to the numbers slice
				num, err := strconv.Atoi(numStr)
				if err == nil {
					numbers = append(numbers, num)
				}
				// Reset the number string for the next number
				numStr = ""
			}
		}
	}

	// Check if there's a pending number string to convert
	if numStr != "" {
		// Convert the number string to an integer and append it to the numbers slice
		num, err := strconv.Atoi(numStr)
		if err == nil {
			numbers = append(numbers, num)
		}
	}
	for _,value := range numbers{
		fmt.Println(value)
	}
	// Print the extracted numbers
	fmt.Println("Extracted numbers:", numbers)
}
