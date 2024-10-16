package main

func LinesPackages(c rune, Lines []string) []string {

	start, end := Limits(c)
	var result []string

	for i := start - 1; i <= end-1; i++ {
		result = append(result, Lines[i])
	}

	return result
}
