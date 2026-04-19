package main

import (
	"strings"
	"unicode"
)

func fixPunctuation(text string) string {
	var out strings.Builder
	runes := []rune(text)

	for i := 0; i < len(runes); i++ {
		r := runes[i]

		if isPunctuation(r) {
			tmp := strings.TrimRight(out.String(), " ")
			out.Reset()
			out.WriteString(tmp)

			out.WriteRune(r)

			for i+1 < len(runes) && isPunctuation(runes[i+1]) {
				i++
				out.WriteRune(runes[i])
			}

			if i+1 < len(runes) && !unicode.IsSpace(runes[i+1]) && !isPunctuation(runes[i+1]) {
				out.WriteRune(' ')
			}

			continue
		}

		out.WriteRune(r)
	}

	return strings.TrimSpace(out.String())
}

func isPunctuation(r rune) bool {
	return strings.ContainsRune(".,!?;:", r)
}
