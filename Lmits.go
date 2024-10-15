package main

func Limits(c rune) (start, end int) {

	n := int(c - ' ')
	start = 9*n + 2
	end = start + 7

	return start, end
}
