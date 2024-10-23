package main

func IsAllNewLines(slice []string) bool {
	for _, each := range slice {
		if each != "\n" {
			return false
		}
	}
	return true
}
