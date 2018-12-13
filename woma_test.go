package woma

import (
	"math"
	"testing"
)

func TestAffixes(t *testing.T) {
	if len(Affixes("")) != 0 {
		t.Fatal("The empty word should have no affixes")
	}
	affixes := Affixes("works")
	if len(affixes) != 8 {
		t.Fatal("Should get 8 affixes from 'works'")
	}
	expected := []string{"w", "wo", "wor", "work",
		"orks", "rks", "ks", "s"}
	for _, e := range expected {
		if !affixes[e] {
			t.Fatal("Could not find", e)
		}
	}
}

func TestNgrams(t *testing.T) {
	if len(Ngrams(2)("")) != 0 {
		t.Fatal("The empty word should have no n-grams")
	}
	bigrams := Ngrams(2)("night")
	if len(bigrams) != 4 {
		t.Fatal("Should get 4 bigrams from 'night'")
	}
	expected := []string{"ni", "ig", "gh", "ht"}
	for _, e := range expected {
		if !bigrams[e] {
			t.Fatal("Could not find", e)
		}
	}
}

func TestDice(t *testing.T) {
	dsc := Dice(Ngrams(2))("night", "nacht")
	if math.Abs(dsc-0.25) > 1e-16 {
		t.Fatal("Should get a Dice score of 0.25")
	}
	dsc = Dice(Affixes)("work", "wirk")
	if math.Abs(dsc-0.5) > 1e-16 {
		t.Fatal("Should get a Dice score of 0.5")
	}
	dsc = Dice(Affixes)("", "wirk")
	if math.Abs(dsc) > 1e-16 {
		t.Fatal("Should get a Dice score of 0")
	}
}
