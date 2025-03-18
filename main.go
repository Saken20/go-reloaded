package main

import (
	"fmt"
	"go-reloaded/fileio"
	"go-reloaded/textprocessor"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run . <input_file> <output_file>")
		return
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	text, err := fileio.ReadFile(inputFile)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	modifiedText := textprocessor.ProcessText(text)

	err = fileio.WriteFile(outputFile, modifiedText)
	if err != nil {
		fmt.Println("Error writing file:", err)
	}
}
