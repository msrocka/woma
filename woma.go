package woma

import (
	"strings"
)

// Normalize converts the given string to lower case without leading and
// trailing whitespaces.
func Normalize(s string) string {
	return strings.ToLower(strings.TrimSpace(s))
}

// Affixes returns the set of prefixes and suffixes of the given word.
func Affixes(word string) map[string]bool {
	nword := Normalize(word)
	affixes := make(map[string]bool)
	if nword == "" {
		return affixes
	}
	runes := []rune(nword)
	if len(runes) < 2 {
		return affixes
	}
	for i := range runes {
		if i == 0 {
			continue
		}
		if i < (len(runes)) {
			affixes[string(runes[0:i])] = true
		}
		affixes[string(runes[i:])] = true
	}
	return affixes
}

// Ngrams constructs a function that calculates the n-grams of a word.
func Ngrams(n int) func(string) map[string]bool {
	return func(word string) map[string]bool {
		ngrams := make(map[string]bool)
		nword := Normalize(word)
		if len(nword) < n {
			return ngrams
		}
		runes := []rune(nword)
		for i := range runes {
			ngram := string(runes[i:(i + n)])
			ngrams[ngram] = true
			if (i + n) == len(runes) {
				break
			}
		}
		return ngrams
	}
}

// Dice returns a function that calculates the Dice coefficient
// https://en.wikipedia.org/wiki/S%C3%B8rensen%E2%80%93Dice_coefficient of
// two words: 2 * |A & B| / (|A| + |B|)
// It takes a function as argument that calculates the sets A and B from the
// given words.
func Dice(fn func(string) map[string]bool) func(string, string) float64 {
	return func(wordA, wordB string) float64 {
		a := fn(wordA)
		b := fn(wordB)
		u := 0
		for e := range a {
			if b[e] {
				u++
			}
		}
		return float64(2*u) / float64(len(a)+len(b))
	}
}
