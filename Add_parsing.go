package main

import (
	"fmt"
	"strconv"
	"strings"

)
func removeElement(slice []string, element string) []string {
	result := []string{}

	for _, v := range slice {
		if v != element {
			result = append(result, v)
		}
	}

	return result
}

	
func Answer(question string) (int, bool) {
	operations := map[string]rune{"add": '+', "minus": '-', "divided": '/', "multiplied": '*', "plus": '+'}
	
	if strings.Contains(question,"What is "){
		num1 := strings.Split(question[len("What is "):len(question)-1], " ")
		
	// fmt.Println(num1)
	num := removeElement(num1,"by")
        // fmt.Println(num)
		if len(num) < 3 || len(num)%2 == 0 {
			return 0, false // Invalid input format
		}

	// Initialize result to the first number
	result, err := strconv.Atoi(num[0])
	if err != nil {
		return 0, false
	}
	// Iterate over the slice starting from index 1
	for i := 1; i < len(num); i += 2 {
		
		
		var num1 int
		op := num[i]
        if _, ok := operations[string(op)]; !ok {
    return 0, false
}
		if i+1<len(num){
			num1, err = strconv.Atoi(num[i+1])
		}
		if err != nil {
			return 0, false
		}
		
		// Perform the operation
		switch operations[string(op)] {
		case '+':
			result += num1
		case '-':
			result -= num1
		case '*':
			result *= num1
		case '/':
			if num1 != 0 {
				result /= num1
			} else {
				return 0, false // Division by zero
			}
        default :
        	return 0,false
		}
	}
	return result, true
	}else{
		return 0,false
	}
}

// func main() {
// 	var question string
// 	question ="What is 52 minus -6?"
	
// 	// fmt.Println(slice)
// 	answer, ok := Answer(question)
// 	if ok {
// 		fmt.Printf("Answer: %d\n", answer) // Output: Answer: 15
// 	} else {
// 		fmt.Println("Invalid input")
// 	}
// }
