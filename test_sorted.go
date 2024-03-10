package main

import (
	"fmt"
	"sort"
	"math/big"
	// "strconv"
)

func bigSorting(unsorted []string) []string {
	sorted := make([]*big.Int, 0)
	intVal := new(big.Int)
	for _, value := range unsorted {
		
		intVal, ok := intVal.SetString(value, 10)
		if !ok {
			fmt.Printf("Failed to parse string %s\n", value)
			continue
		}
		sorted = append(sorted, intVal)
	}

	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i].Cmp(sorted[j]) < 0
	})

	sortedStrings := make([]string, 0)
	for _, value := range sorted {
		st := value.String()
		sortedStrings = append(sortedStrings, st)
	}

	return sortedStrings
}
func main() {
	unsorted := []string{"31415926535897932384626433832795", "1", "3", "10", "3", "5"}
	// sorted := make([]*big.Int, 0)

	// for _, value := range unsorted {
	// 	intVal := new(big.Int)
	// 	intVal, ok := intVal.SetString(value, 10)
	// 	if !ok {
	// 		fmt.Printf("Failed to parse string %s\n", value)
	// 		continue
	// 	}
	// 	sorted = append(sorted, intVal)
	// }
	// fmt.Println(sorted)
	sorted := bigSorting(unsorted)
	fmt.Println(sorted)
	a := new(big.Int).SetInt64(1234567890123456789)
	b ,c:= new(big.Int).SetString("98765432109876543210", 10)
	// e ,d:= new(big.Int).SetString("98765432109878543210", 10)
	fmt.Println(b,c,)
	// Perform arithmetic operations
	sum := new(big.Int).Add(a, b)
	diff := new(big.Int).Sub(b, a)
	prod := new(big.Int).Mul(a, b)
	quo := new(big.Int).Div(b, a)

	// Print results
	fmt.Println("Sum:", sum)
	fmt.Println("Difference:", diff)
	fmt.Println("Product:", prod)
	fmt.Println("Quotient:", quo)
}
