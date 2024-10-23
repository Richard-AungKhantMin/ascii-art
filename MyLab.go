package main

func MyLab(input string) string {

	var FinalText string
	var emptyLineCount int

	Filtered := Filter(input)

	if IsAllNewLines(Filtered) {
		for i := 0; i < len(Filtered); i++ {
			FinalText = FinalText + "\n"
		}
		return FinalText
	}

	for i := 0; i < len(Filtered); i++ {
		if Filtered[i] == "\n" {

			if emptyLineCount > 0 {
				FinalText = FinalText + "\n"
			}
			emptyLineCount++

		} else {

			FinalText = FinalText + LinesOfWords(Filtered[i]) + "\n"
		}
	}

	return FinalText
}
