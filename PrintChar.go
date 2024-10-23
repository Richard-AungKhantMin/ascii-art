package main

func PrintChar(word []string) string {

	var resultText string

	for i := 0; i < len(word); i++ {
		resultText = resultText + "\n" + word[i]
	}
	resultText = resultText + "\n"

	return resultText
}
