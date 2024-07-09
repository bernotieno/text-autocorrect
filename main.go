package main

import (
	"bufio"
	"fmt"
	"os"

	"go-reloaded/transform"
)

func main() {
	// Check if the correct number of command line arguments is provided
	if len(os.Args) != 3 {
		fmt.Println("Invalid parameters, usage: go run . par1 par2")
		return
	}

	// Extract command line arguments
	args := os.Args[1:]

	// Open the input file
	file, err := os.Open(args[0])
	if err != nil {
		fmt.Println(err) // Print error if opening file fails
	}
	defer file.Close() // Ensure file is closed after function returns

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	var result string
	// Read lines from file and concatenate them into a single string
	for scanner.Scan() {
		lines := scanner.Text()

		result += transform.TransformWords(lines) + "\n"

	}

	// Check for any scanning errors
	if err := scanner.Err(); err != nil {
		fmt.Println("Error: ", err)
	}

	// Write the final string to the output file
	doc := os.WriteFile(args[1], []byte(result), 0o644)
	if doc != nil {
		panic(doc) // Panic if writing to file fails
	}

	// fmt.Println("File was successfully created and edited.")
}
