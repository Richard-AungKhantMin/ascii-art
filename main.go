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

	style, err := os.Open(os.Args[2])
	IsErrNil(err)
	defer style.Close()

	Lines := bufio.NewScanner(style)
	var NestedSlice [][]string

	for _, eachC := range os.Args[1] {
		NestedSlice = append(NestedSlice, LinesPackage(eachC, Lines))
	}

	for _, eachSlice := range NestedSlice {
		fmt.Println(eachSlice)
	}

}
