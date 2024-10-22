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

	Filtered := Filter(os.Args[1])
	fmt.Println(len(Filtered))

	for i := 0; i < len(Filtered); i++ {
		if Filtered[i] == "\\n" {
			fmt.Println()
		} else {
			MyLab(Filtered[i])
		}
	}

}
