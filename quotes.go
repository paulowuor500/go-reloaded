package main

import "strings"

func fixQuotes(text string) string {
	words := strings.Fields(text)
	var result []string
	opening := true

	for i := 0; i < len(words); i++ {
		if words[i] == "'" {
			if opening {
				if i+1 < len(words) {
					words[i+1] = "'" + words[i+1]
				}
				opening = false
			} else {
				if len(result) > 0 {
					result[len(result)-1] = result[len(result)-1] + "'"
				}
				opening = true
			}
		} else {
			result = append(result, words[i])
		}
	}
	return strings.Join(result, " ")
}
