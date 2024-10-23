package main

import "testing"

func TestMyLab(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected string
	}{
		{"Hello", "Hello\n", ` _    _          _   _          
		| |  | |        | | | |         
		| |__| |   ___  | | | |   ___   
		|  __  |  / _ \ | | | |  / _ \  
		| |  | | |  __/ | | | | | (_) | 
		|_|  |_|  \___| |_| |_|  \___/  
										
										
		`},
	}

	// Iterate through each test case
	for _, tc := range testCases {
		// Run the function with the test input
		result := MyLab(tc.input)

		// Compare the result with the expected output
		if result != tc.expected {
			t.Errorf("MyLab(%q) = %q; expected %q", tc.input, result, tc.expected)
		}
	}
}
