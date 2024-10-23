package main

func Mylab(input string) string {

	var emptyLineCount int
	var FinalText string
	Filtered := Filter(input)

	for i := 0; i < len(Filtered); i++ {
		if Filtered[i] == "\n" {

			if emptyLineCount > 0 {
				FinalText = FinalText + "\n"
			}
			emptyLineCount++

		} else {

			FinalText = FinalText + LineOfWords(Filtered[i])
		}
	}

	return FinalText
}
