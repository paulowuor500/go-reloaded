package main

import (
	"strings"
	"unicode"
)

func applyArticles(words []string) []string {
	for i := 0; i < len(words)-1; i++ {
		if isArticleA(words[i]) && startsWithVowelOrH(words[i+1]) {
			if words[i] == "A" {
				words[i] = "An"
			} else {
				words[i] = "an"
			}
		}
	}
	return words
}

func isArticleA(w string) bool {
	return w == "a" || w == "A"
}

func startsWithVowelOrH(word string) bool {
	for _, r := range word {
		if unicode.IsLetter(r) {
			r = unicode.ToLower(r)
			return strings.ContainsRune("aeiouh", r)
		}
	}
	return false
}
