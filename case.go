package main

import (
	"strconv"
	"strings"
	"unicode"
)

func handleCase(words []string) []string {
	var result []string

	for i := 0; i < len(words); i++ {
		word := words[i]

		// Check for multi-word tags like (up, 2)
		// This happens if the current word starts with (up, and there is a next word
		if strings.HasPrefix(word, "(") && strings.HasSuffix(word, ",") && i+1 < len(words) {
			mode := strings.Trim(word, "(,")
			countStr := strings.Trim(words[i+1], ")")
			count, err := strconv.Atoi(countStr)
			
			if err == nil {
				applyCase(result, len(result)-1, count, "("+mode+")")
				i++ // Skip the number word
				continue
			}
		}

		// Check for single word tags like (up)
		mode, count, ok := parseCase(word)
		if ok && count == 1 {
			applyCase(result, len(result)-1, 1, mode)
			continue
		}

		result = append(result, word)
	}
	return result
}

//robust parser (fixes your bug)
func parseCase(word string) (string, int, bool) {
	// clean punctuation
	clean := ""
	for _, r := range word {
		if unicode.IsLetter(r) || unicode.IsNumber(r) || r == ',' || r == '(' || r == ')' {
			clean += string(r)
		}
	}

	clean = strings.Trim(clean, "()")

	// CASE 1: simple (up), (low), (cap)
	if clean == "up" || clean == "low" || clean == "cap" {
		return "(" + clean + ")", 1, true
	}

	// CASE 2: (up, N)
	parts := strings.Split(clean, ",")
	if len(parts) != 2 {
		return "", 0, false
	}

	mode := parts[0]
	count, err := strconv.Atoi(parts[1])
	if err != nil || count <= 0 {
		return "", 0, false
	}

	switch mode {
	case "up", "low", "cap":
		return "(" + mode + ")", count, true
	}

	return "", 0, false
}

func applyCase(words []string, end int, count int, mode string) {
	start := end - count + 1
	if start < 0 {
		start = 0
	}

	for i := start; i <= end; i++ {
		switch mode {
		case "(up)":
			words[i] = strings.ToUpper(words[i])
		case "(low)":
			words[i] = strings.ToLower(words[i])
		case "(cap)":
			words[i] = capitalize(words[i])
		}
	}
}

func capitalize(s string) string {
	if len(s) == 0 {
		return s
	}
	return strings.ToUpper(string(s[0])) + strings.ToLower(string(s[1:]))
}
