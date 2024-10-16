package main

import "bufio"

func LinesPackage(c rune, Lines *bufio.Scanner) (result [8]string) {

	lineCounter := 1
	start, end := Limits(c)

	for Lines.Scan() {

		if lineCounter == start {
			for i := 0; i < len(result); i++ {

				if lineCounter <= end {
					result[i] = Lines.Text()
					lineCounter++
				}

			}

		}

		if lineCounter > end {
			break
		}

		lineCounter++
	}

	return result
}
