package main

import (
	"strconv"
	"strings"
)

func handleConversions(words []string) []string {
	var result []string

	for i := 0; i < len(words); i++ {
		w := strings.ToLower(words[i])

		if (w == "(hex)" || w == "(bin)") && len(result) > 0 {
			prev := result[len(result)-1]

			num, punct := splitNumberAndPunct(prev)

			var converted string
			if w == "(hex)" {
				converted = hexToDecimal(num)
			} else {
				converted = binToDecimal(num)
			}

			result[len(result)-1] = converted + punct
			continue
		}

		result = append(result, words[i])
	}

	return result
}

func splitNumberAndPunct(s string) (string, string) {
	i := len(s)
	for i > 0 && strings.ContainsRune(".,!?;:", rune(s[i-1])) {
		i--
	}
	return s[:i], s[i:]
}

func hexToDecimal(s string) string {
	n, err := strconv.ParseInt(s, 16, 64)
	if err != nil {
		return s
	}
	return strconv.Itoa(int(n))
}

func binToDecimal(s string) string {
	n, err := strconv.ParseInt(s, 2, 64)
	if err != nil {
		return s
	}
	return strconv.Itoa(int(n))
}
