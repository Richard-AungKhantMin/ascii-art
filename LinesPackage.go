package main

import "bufio"

func LinesPackages(c rune, Lines *bufio.Scanner) (Finalresult []string) {

	lineCounter := 1
	start, end := Limits(c)
	var result [8]string

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

	Finalresult = result[:]

	return Finalresult
}
