package main

import (
	"bufio"
	"fmt"
	"os"
)

func BufioToSlice() []string {
	style, err := os.Open("standard.txt")
	IsErrNil(err)
	defer style.Close()

	var Lines []string
	TempoForReading := bufio.NewScanner(style)

	for TempoForReading.Scan() {
		Lines = append(Lines, TempoForReading.Text())
	}

	if err := TempoForReading.Err(); err != nil {
		fmt.Println("Error reading lines:", err)
	}

	return Lines

}
