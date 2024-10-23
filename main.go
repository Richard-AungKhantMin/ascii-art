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
	var UltraMegaFinalResultText string
	Filtered := Filter(os.Args[1])

	for i := 0; i < len(Filtered); i++ {
		if Filtered[i] == "\n" {

			if emptyLineCount > 0 {
				UltraMegaFinalResultText = UltraMegaFinalResultText + "\n"
			}
			emptyLineCount++

		} else {
			fmt.Println(Filtered[i])
			UltraMegaFinalResultText = UltraMegaFinalResultText + MyLab(Filtered[i])
		}
	}

	fmt.Println(UltraMegaFinalResultText)
}
