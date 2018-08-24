package woma

import (
	"testing"
)

func TestParse(t *testing.T) {
	phrase := Parse("Live palometa fish, or mylossoma aureum")
	if len(phrase) != 6 {
		t.Fatal("The phrase should have 6 words")
	}
}
