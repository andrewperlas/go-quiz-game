package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	// Open the CSV file
	file, err := os.Open("problems.csv")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	// Create a new CSV reader
	reader := csv.NewReader(file)

	// Read all records from CSV file
	for {
		// Read the next record
		record, err := reader.Read()
		// If we reached the end of the file, break the loop
		if err != nil {
			if err.Error() == "End of File" {
				break
			}

			fmt.Println("Error", err)
			return
		}

		// Print question 1, which is the first column of the first line
		fmt.Println("Question:", record[0], " =")

		// User inputs an answer and program compares input with the value of the second column
		var answer string
		fmt.Scanln(&answer)
	}

}
