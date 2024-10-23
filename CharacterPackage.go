package main

func CharacterPackage(c rune, Lines []string) []string {

	start, end := Limits(c)
	var result []string

	for i := start; i <= end; i++ {
		result = append(result, Lines[i])
	}

	return result
}
