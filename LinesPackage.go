package main

import (
	"bufio"
)

func LinesPackages(c rune, Lines *bufio.Scanner) []string {

	lineCounter := 1
	start, end := Limits(c)
	var result []string

	for Lines.Scan() {

		if lineCounter >= start && lineCounter <= end {
			result = append(result, Lines.Text())
		}

		if lineCounter > end {
			break
		}

		lineCounter++
	}

	return result
}
