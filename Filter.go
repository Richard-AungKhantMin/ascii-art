package main

func Filter(sentence string) []string {
	var FilteredSlice []string
	var temp string

	for i := 0; i < len(sentence); i++ {
		// Check for '\n'
		if i < len(sentence)-1 && sentence[i] == '\\' && sentence[i+1] == 'n' {
			// Add any accumulated string before the '\n'
			if temp != "" {
				FilteredSlice = append(FilteredSlice, temp)
				temp = "" // Reset temp
			}
			// Add the newline character to the slice
			FilteredSlice = append(FilteredSlice, "\n")
			i++ // Skip the 'n' character as we already processed '\n'
		} else {
			// Accumulate normal characters into temp
			temp += string(sentence[i])
		}
	}

	// Add any remaining string to the slice
	if temp != "" {
		FilteredSlice = append(FilteredSlice, temp)
	}

	return FilteredSlice
}
