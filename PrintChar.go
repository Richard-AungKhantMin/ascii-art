package main

func PrintChar(word []string) string {

	resultText := word[0]

	for i := 1; i < len(word); i++ {
		resultText = resultText + "\n" + word[i]
	}
	resultText = resultText + "\n"

	return resultText
}
