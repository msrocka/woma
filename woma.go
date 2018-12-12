package woma

import (
	"strings"
)

// Normalize converts the given string to lower case without leading and
// trailing whitespaces.
func Normalize(s string) string {
	return strings.ToLower(strings.TrimSpace(s))
}

func Ngrams(word string) []string {
	nword := Normalize(word)
	var ngrams []string
	if nword == "" {
		return ngrams
	}
	runes := []rune(nword)
	if len(runes) < 2 {
		return ngrams
	}
	for i := range runes {
		if i == 0 {
			continue
		}
		if i < (len(runes)) {
			ngrams = append(ngrams, string(runes[0:i]))
		}
		ngrams = append(ngrams, string(runes[i:]))
	}
	return ngrams
}
