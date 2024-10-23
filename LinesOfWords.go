package main

func LinesOfWords(word string) string {
	Lines := BufioToSlice()

	var NestedSlice [][]string
	for _, eachC := range word {
		NestedSlice = append(NestedSlice, CharacterPackage(eachC, Lines))
	}

	FinalResult := make([]string, 8)
	for i := 0; i <= 7; i++ {
		for j := 0; j < len(NestedSlice); j++ {
			FinalResult[i] = FinalResult[i] + NestedSlice[j][i]
		}
	}

	return PrintChar((FinalResult))
}
