package main

import (
	"fmt"
	"log"
	"os"
	"go-reloaded/utilities"
)

func main() {
	// Check if correct number of arguments are provided
	if len(os.Args) != 3 {
		fmt.Println("Usage: main input_file output_file")
		return
	}

	// Read input file
	inputFileName := os.Args[1]
	inputData, err := os.ReadFile(inputFileName)
	if err != nil {
		log.Fatalf("Error reading input file: %v", err)
	}

	// Apply modifications
	modifiedData := utilities.LowerCase(string(inputData))
	modifiedData = utilities.UppercaseComplete(modifiedData)
	modifiedData = utilities.AtoAn(modifiedData)
	modifiedData = utilities.HexToDecimal(modifiedData)
	modifiedData = utilities.CapitalizePrevious(modifiedData)
	modifiedData = utilities.BinarytoDecimal(modifiedData)
	modifiedData = utilities.FormatText(modifiedData)
	modifiedData = utilities.Quotation(modifiedData)

	// Write modified data to output file
	outputFileName := os.Args[2]
	err = os.WriteFile(outputFileName, []byte(modifiedData), 0o644)
	if err != nil {
		log.Fatalf("Error writing output file: %v", err)
	}

	fmt.Println("Modifications applied successfully.")
}
