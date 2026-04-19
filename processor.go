package main

import "strings"

func ProcessText(text string) string {
	lines := strings.Split(text, "\n")

	for i, line := range lines {
		if strings.TrimSpace(line) == "" {
			lines[i] = ""
			continue
		}

		words := strings.Fields(line)

		words = handleConversions(words)
		words = handleCase(words)

		// join BEFORE punctuation/quotes/articles
		result := strings.Join(words, " ")

		result = fixPunctuation(result)
		result = fixQuotes(result)

		// articles work on final string (NOT []string)
		words = strings.Fields(result)
		words = applyArticles(words)

		lines[i] = strings.Join(words, " ")
	}

	return strings.Join(lines, "\n")
}
