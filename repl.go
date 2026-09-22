package main

import (
	"strings"
)

func cleanInput(text string) []string {
	// Split the string into "words" my whitespace
	// Lowercase and trim leading/trailing whitespace

	// Replace any newline/ tab characters with spaces for splitting
	text = strings.Replace(text, "\n", " ", -1)
	text = strings.Replace(text, "\t", " ", -1)

	textTrim := strings.TrimSpace(text)
	textLower := strings.ToLower(textTrim)
	textSplit := strings.Split(textLower, " ")

	var cleanString []string
	for _, v := range textSplit {
		trim := strings.TrimSpace(v)
		if trim == "" {
			continue
		}
		cleanString = append(cleanString, trim)
	}

	// fmt.Printf("Clean input:\t%q, Length:%v\n", cleanString, len(cleanString))

	return cleanString

}
