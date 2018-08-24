package woma

import (
	"bytes"
	"unicode"
	"unicode/utf8"
)

// A Phrase in woma is just a list of (normalized) words (without punctuation etc.)
type Phrase []string

// Parse collects the words from the given text into a phrase.
func Parse(text string) Phrase {
	var phrase Phrase
	var word bytes.Buffer
	addWord := func() {
		if word.Len() == 0 {
			return
		}
		w := Normalize(word.String())
		word.Reset()
		if utf8.RuneCountInString(w) == 0 {
			return
		}
		phrase = append(phrase, w)
	}
	for _, char := range text {
		if !isWordChar(char) {
			addWord()
			continue
		}
		word.WriteRune(char)
	}
	addWord()
	return phrase
}

func isWordChar(char rune) bool {
	if unicode.IsLetter(char) {
		return true
	}
	if unicode.IsDigit(char) {
		return true
	}
	if char == '-' {
		return true
	}
	return false
}
