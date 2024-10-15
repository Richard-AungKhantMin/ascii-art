package main

func Limits(c rune) (start, end int) {

	n := int(c-' ') + 1
	start = 9*n + 1
	end = start + 7

	return start, end
}
