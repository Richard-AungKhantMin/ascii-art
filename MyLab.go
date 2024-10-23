package main

import "strings"

func Mylab(input string) string {

	var emptyLineCount int
	var FinalText string
	Splited := strings.Split(input, "\\n")

	for i := 0; i < len(Splited); i++ {
		if Splited[i] == "\n" {

			if emptyLineCount > 0 {
				FinalText = FinalText + "\n"
			}
			emptyLineCount++

		} else {

			FinalText = FinalText + LineOfWords(Splited[i])
		}
	}

	return FinalText
}
