package main

import (
	"strings"
)

func MyLab(input string) string {

	var FinalText string

	Splited := strings.Split(input, "\\n")

	for i := 0; i < len(Splited); i++ {
		if Splited[i] == "" {

			FinalText = FinalText + "\n"

		} else {

			FinalText = FinalText + LinesOfWords(Splited[i])
		}
	}

	return FinalText + "\n"
}
