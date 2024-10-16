//1+8+1+8+1+8+1+8+
//1+8 format

// unprintable character
package main

import (
	"bufio"
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

	Lines := bufio.NewScanner(style)
	var NestedSlice [][]string

	for _, eachC := range os.Args[1] {
		NestedSlice = append(NestedSlice, LinesPackages(eachC, Lines))
	}

	for _, eachSlice := range NestedSlice {
		PrintChar(eachSlice)
	}

}
