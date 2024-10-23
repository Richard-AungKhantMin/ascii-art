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
	TempForReading := bufio.NewScanner(style)

	for TempForReading.Scan() {
		Lines = append(Lines, TempForReading.Text())
	}

	if err := TempForReading.Err(); err != nil {
		fmt.Println("Error reading lines:", err)
	}

	return Lines

}
