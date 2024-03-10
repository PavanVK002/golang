package main

import (
	"fmt"
	"strings"
)

func main() {
	var newStr string 
	newStr =strings.ReplaceAll("3-598-21508-8","-","")
	fmt.Printf("%s\n",newStr)
	str := "Hello, 世界" // String containing English and Chinese characters
	fmt.Println("Length of string:", len(str)) // Prints 13 (bytes)
	// Convert the string to a slice of runes
	runes := []rune(str)

	// Get the length of the slice of runes
	length := len(runes)
	fmt.Println("Length of string:", length)
	// Iterating over the string and printing each byte
	for i := 0; i < len(str); i++ {
		fmt.Printf("Byte at index %d: %x\n", i, str[i])
	}

	// Iterating over the string and printing each rune
	for _, r := range str {
		fmt.Printf("Rune: %c\n", r)
	}
}
