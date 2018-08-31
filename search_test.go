package woma

import (
	"testing"
)

func TestSearchListFind(t *testing.T) {
	list := NewSearchList([]string{
		"Weather stripping",
		"Thermal insulation kits",
		"Acoustical insulation",
		"Fireproofing material"})

	_, text := list.Find("Fiber insulation")
	if text != "Acoustical insulation" {
		t.Fatal("should match `Acoustical insulation`")
	}

	_, text = list.Find("Thermal insulation kits")
	if text != "Thermal insulation kits" {
		t.Fatal("should match exactly")
	}

	_, text = list.Find("Fiber insulation kits")
	if text != "Thermal insulation kits" {
		t.Fatal("should match exactly")
	}
}
