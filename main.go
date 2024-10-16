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

	TempoForReading := bufio.NewScanner(style)
	var Lines []string

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
		fmt.Println(string(eachC))
	}

	fmt.Println(NestedSlice[0])

	for _, eachSlice := range NestedSlice {
		PrintChar(eachSlice)
	}

}
