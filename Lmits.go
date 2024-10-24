package main

import (
	"fmt"
	"os"
)

func Limits(c rune) (start, end int) {

	if c <= ' ' || c >= '~' {
		fmt.Println("The input you just put is not a printable character or a character outside of ascii table")
		os.Exit(1)
	}

	n := int(c - ' ')
	start = 9*n + 1
	end = start + 7

	return start, end
}
