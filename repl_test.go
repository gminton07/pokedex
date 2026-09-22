package main

import (
	"fmt"
	"testing"
)

func TestCleanInput(t *testing.T) {
	// Define test cases
	cases := []struct {
		input 	string
		expected 	[]string
	}{
		{
			input: "   hello  world   ",
			expected: []string{"hello", "world"},
		},
		{
			input: "Charmander Bulbasaur PIKACHU",
			expected: []string{"charmander", "bulbasaur", "pikachu"},
		},
		{
			input: "  ",
			expected: []string{},
		},
		{
			input: "Sentences separated \n by a newline",
			expected: []string{"sentences", "separated", "by", "a", "newline"},
		},
		{
			input: "Newlines without\nspaces",
			expected: []string{"newlines", "without", "spaces"},
		},
		{
			input: "TAB\tCHARACTERS",
			expected: []string{"tab", "characters"},
		},
		// add more cases
	}

	// Loop over cases and run tests
	for _, c := range cases {
	actual := cleanInput(c.input)

	// Print actual/expected values
	fmt.Printf("Actual:\t\t%q\n", actual)
	fmt.Printf("Expected:\t%q\n", c.expected)

	// Check the length of the actual slice
	// if they don't match, use t.Errorf and continue to the next case
	
	if len(actual) != len(c.expected) {
		// error and continue here
		t.Errorf("Test fail: Inconsistent lengths\nActual:\t\t%d\nExpected:\t%d", len(actual), len(c.expected))
	}
	for i := range actual {
		word := actual[i]
		expectedWord := c.expected[i]
		// Check each word in the slice
		// if they don't match, use t.Errorf to print an error message
		if word != expectedWord {
			t.Errorf("Test fail: Word mismatch\nActual:\t\t%q\nExpected:\t%q", word, expectedWord)
		}
		// and fail the test
	}
}
}
