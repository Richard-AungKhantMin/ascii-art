package main

import (
	"fmt"
	"os"
)

func main() {

	if len(os.Args) != 2 {
		return
	}

	if os.Args[1] == "" {
		fmt.Println()
		return
	}

	var emptyLineCount int
	Filtered := Filter(os.Args[1])

	for i := 0; i < len(Filtered); i++ {
		if Filtered[i] == "\n" {

			if emptyLineCount > 0 {
				fmt.Println()
			}
			emptyLineCount++

		} else {
			MyLab(Filtered[i])
		}
	}

}
