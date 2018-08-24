package woma

import (
	"testing"
)

func TestParse(t *testing.T) {
	phrase := Parse("Live palometa fish, or mylossoma aureum")
	if len(phrase) != 6 {
		t.Fatal("The phrase should have 6 words")
	}
	words := []string{"live", "palometa", "fish", "or", "mylossoma", "aureum"}
	for i := range words {
		if phrase[i] != words[i] {
			t.Fatal("Word", phrase[i], "!=", words[i])
		}
	}
}

func TestWithoutStopwords(t *testing.T) {
	phrase := Parse("Live palometa fish, or mylossoma aureum").WithoutStopwords()
	if len(phrase) != 5 {
		t.Fatal("The phrase should have 5 words")
	}
	words := []string{"live", "palometa", "fish", "mylossoma", "aureum"}
	for i := range words {
		if phrase[i] != words[i] {
			t.Fatal("Word", phrase[i], "!=", words[i])
		}
	}
}
