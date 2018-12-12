package woma

import (
	"testing"
)

func TestNgrams(t *testing.T) {
	if len(Ngrams("")) != 0 {
		t.Fatal("The empty word should have no n-grams")
	}
	ngrams := Ngrams("works")
	if len(ngrams) != 8 {
		t.Fatal("Should get 8 n-grams from 'works'")
	}
	expected := []string{"w", "wo", "wor", "work",
		"orks", "rks", "ks", "s"}
	for _, e := range expected {
		for i, ng := range ngrams {
			if e == ng {
				break
			}
			if i == (len(ngrams) - 1) {
				t.Fatal("Could not find", e)
			}
		}
	}
}
