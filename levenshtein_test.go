package woma

import (
	"testing"
)

func TestLevenshtein(t *testing.T) {
	if Levenshtein("", "") != 0 {
		t.Fatal("lev(,) != 0")
	}
	if Levenshtein("über", "über") != 0 {
		t.Fatal("lev(,) != 0")
	}
	if Levenshtein("acb", "adb") != 1 {
		t.Fatal("lev(acb, adb) != 1")
	}
	if Levenshtein("über", "übäer") != 1 {
		t.Fatal("lev(über, übäer) != 1")
	}
	if Levenshtein("über", "übäeer") != 2 {
		t.Fatal("lev(über, übäeer) != 2")
	}
	if Levenshtein("trüber", "träbäeer") != 3 {
		t.Fatal("lev(trüber, träbäeer) != 3")
	}
}
