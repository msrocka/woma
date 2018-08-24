package woma

import (
	"testing"
)

func TestLevenshtein(t *testing.T) {
	if Levenshtein("acb", "adb") != 1 {
		t.Fatal("lev(acb, adb) != 1")
	}
	if Levenshtein("über", "übäer") != 2 {
		t.Fatal("lev(über, übäer) != 2")
	}
}
