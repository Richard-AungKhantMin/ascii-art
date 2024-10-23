package main

func Filter(sentence string) []string {

	var FilteredSlice []string
	var word string
	lastWord := len(sentence) - 1

	for i := 0; i <= lastWord; i++ {

		if i < lastWord && sentence[i] == '\\' && sentence[i+1] == 'n' {

			if word != "" {
				FilteredSlice = append(FilteredSlice, word)
				word = ""
			}

			FilteredSlice = append(FilteredSlice, "\n")
			i++

		} else {

			word = word + string(sentence[i])
		}
	}

	if word != "" {
		FilteredSlice = append(FilteredSlice, word)
	}

	return FilteredSlice
}
