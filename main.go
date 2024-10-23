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

	UltraMegaFinalResultText := Mylab(os.Args[1])
	fmt.Print(UltraMegaFinalResultText)
}
