package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"reflect"
	"os"
)

func main() {
	// Open the input CSV file
	inputFile, err := os.Open("Security_codes.csv")
	if err != nil {
		log.Fatalf("Error opening input file: %v", err)
	}
	defer inputFile.Close()

	// Create a CSV reader for the input file
	reader := csv.NewReader(inputFile)
	fmt.Println(reflect.TypeOf(reader))
	// Read all records from the input CSV
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatalf("Error reading input CSV: %v", err)
	}

	// Create a new CSV file for output
	// outputFile, err := os.Create("Security_codes.csv")
	// if err != nil {
	// 	log.Fatalf("Error creating output file: %v", err)
	// }
	// defer outputFile.Close()
	
	// // Create a CSV writer for the output file
	// writer := csv.NewWriter(outputFile)
	// defer writer.Flush()

	// Initialize a slice to store the six-digit codes
	// var codes []string
	fmt.Println(reflect.TypeOf(records),len(records))
	// // Extract the six-digit codes from each record and append to the codes slice
	// for _, record := range records {
	// 	 // Print progress
	// 	if len(record) > 0 {
	// 		codes = append(codes, record[0])
	// 	}
	// }

	// Write the codes slice to the output CSV file
	// for _, code := range codes {
	// 	err := writer.Write([]string{code})
	// 	if err != nil {
	// 		log.Fatalf("Error writing to output CSV: %v", err)
	// 	}
	// }

	fmt.Println("Extracted six-digit codes saved to output.csv")
}
