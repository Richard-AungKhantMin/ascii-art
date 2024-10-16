//1+8+1+8+1+8+1+8+
//1+8 format

// unprintable character
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

	style, err := os.Open(os.Args[2])
	IsErrNil(err)
	defer style.Close()

	var Lines []string
	TempoForReading := bufio.NewScanner(style)

	for TempoForReading.Scan() {
		Lines = append(Lines, TempoForReading.Text())
	}

	if err := TempoForReading.Err(); err != nil {
		fmt.Println("Error reading lines:", err)
		return
	}

	var NestedSlice [][]string
	for _, eachC := range os.Args[1] {
		NestedSlice = append(NestedSlice, LinesPackages(eachC, Lines))
	}

	FinalResult := make([]string, 8)
	for i := 0; i <= 7; i++ {
		for j := 0; j < len(NestedSlice); j++ {
			FinalResult[i] = FinalResult[i] + NestedSlice[j][i]
		}
	}

	PrintChar(FinalResult)
}
