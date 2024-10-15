//1+8+1+8+1+8+1+8+
//1+8 format

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	if len(os.Args) != 3 {
		return
	}

	if os.Args[1] == "" {
		return
	}

	// start coding here.

	// Open the file
	file1, err1 := os.Open(os.Args[2])
	if err1 != nil {
		fmt.Println("Error opening file1:", err1)
		return
	}
	// Ensure the file1 is closed after function completes
	defer file1.Close()

	// Create a scanner to read file1 line by line
	scanner1 := bufio.NewScanner(file1)

	// Define the range of lines to combine (for example, lines 5 to 10)
	startLine := 5
	endLine := 10
	lineCounter := 1    // Initialize line counter
	combinedLines := "" // String to store the combined lines

	// Read file1 line by line
	for scanner1.Scan() {
		if lineCounter >= startLine && lineCounter <= endLine {
			line1 := scanner1.Text()     // Get the current line from file1
			combinedLines += line1 + " " // Combine lines with a space (or any other delimiter)
		}
		if lineCounter > endLine {
			break // Stop after processing the specified range of lines
		}
		lineCounter++ // Increment the line counter
	}

	// Print the combined lines
	fmt.Println("Combined lines:", combinedLines)

	// Check for errors during scanning of file1
	if err2 := scanner1.Err(); err2 != nil {
		fmt.Println("Error reading file1:", err2)
	}
}
